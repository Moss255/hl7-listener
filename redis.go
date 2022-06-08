package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/lenaten/hl7"
)

func forwardMessageREDIS(msg *hl7.Message, url string) error {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	err := rdb.LPush(ctx, "HL7", string(msg.Value)).Err()

	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}
