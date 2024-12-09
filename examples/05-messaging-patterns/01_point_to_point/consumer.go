package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "point-to-point-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	consumer.SubscribeTopics([]string{"point-to-point-topic"}, nil)
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Println("Error reading message:", err)
			continue
		}
		fmt.Println("Received:", string(msg.Value))
	}
}
