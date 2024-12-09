package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "service-a",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	consumer.SubscribeTopics([]string{"message-bus-topic"}, nil)
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("Service A received:", string(msg.Value))
	}
}
