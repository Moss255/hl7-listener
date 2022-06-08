package main

import (
	"fmt"

	"github.com/lenaten/hl7"
	"github.com/streadway/amqp"
)

func ForwardToRabbitQueue(msg *hl7.Message, url string) {

	conn, err := amqp.Dial(url)

	if err != nil {
		fmt.Print(err)
	}

	defer conn.Close()

	channel, err := conn.Channel()

	if err != nil {
		fmt.Print(err)
	}

	defer channel.Close()

	queue, err := channel.QueueDeclare("Test", true, false, false, false, nil)

	if err != nil {
		fmt.Print(err)
	}

	body := string(msg.Value)

	err = channel.Publish("", queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})

	if err != nil {
		fmt.Print(err)
	}
}
