package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "publish-subscribe-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	consumer.SubscribeTopics([]string{"publish-subscribe-topic"}, nil)
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Println("Error reading message:", err)
			continue
		}
		fmt.Println("Subscriber received:", string(msg.Value))
	}
}
