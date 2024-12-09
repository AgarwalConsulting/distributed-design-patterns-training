package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	topic := "point-to-point-topic"
	message := "Task for one consumer"

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)
	if err != nil {
		fmt.Println("Failed to produce message:", err)
		return
	}
	fmt.Println("Message sent:", message)
	producer.Flush(1000)
}
