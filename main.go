package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lenaten/hl7"
)

func forwardMessageHTTP(msg *hl7.Message, url string) error {

	client := http.Client{}

	body := bytes.NewReader(msg.Value)

	_, err := client.Post(url, "text/plain", body)

	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}

func sendAcknowledgement(msg *hl7.Message) (*hl7.Message, error) {

	info, err := msg.Info()

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	ack := hl7.Acknowledge(info, nil)

	msg3 := hl7.NewLocation("MSH.3")
	msg4 := hl7.NewLocation("MSH.4")

	ack.Set(msg3, "SWLSTG HL7 Reciever")
	ack.Set(msg4, "SWLSTG HL7 Reciever")

	return ack, nil
}

func HL7Handler(res http.ResponseWriter, req *http.Request) {

	fmt.Printf("Incoming request from %s\n", req.RemoteAddr)

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)

	if err != nil {
		fmt.Print(err)
	}

	// Strange Quirk with golang
	body = bytes.Trim(body, "EOF")

	msg := hl7.NewMessage(body)

	url := os.Getenv("FORWARDING_ADDRESS")

	//msgErr := forwardMessageREDIS(msg, url)

	ForwardToRabbitQueue(msg, url)

	//if msgErr != nil {
	//	fmt.Print(msgErr)
	//}

	ack, err := sendAcknowledgement(msg)

	res.Write(ack.Value)
}

func main() {

	os.Setenv("FORWARDING_ADDRESS", "http://localhost:8081")

	http.HandleFunc("/", HL7Handler)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	fmt.Printf("Listening on port %s\n", PORT)
	fmt.Printf("Incoming messages will be forwarded onto %s", os.Getenv("FORWARDING_ADDRESS"))

	addr := fmt.Sprintf(":%s", PORT)

	http.ListenAndServe(addr, nil)
}
