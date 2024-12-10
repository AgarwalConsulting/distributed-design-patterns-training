package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./messages.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create table to store processed messages
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS processed_messages (message_id TEXT PRIMARY KEY, content TEXT)")
	if err != nil {
		log.Fatal(err)
	}
}

func processMessage(messageID, content string) {
	// Check if message is already processed
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM processed_messages WHERE message_id = ?)", messageID).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		fmt.Println("Message already processed, skipping:", messageID)
		return
	}

	// Process and store message
	_, err = db.Exec("INSERT INTO processed_messages (message_id, content) VALUES (?, ?)", messageID, content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Processed message:", messageID)
}

func main() {
	// Connect to Kafka
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	// Subscribe to a Kafka topic
	partitionConsumer, err := consumer.ConsumePartition("example-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal(err)
	}
	defer partitionConsumer.Close()

	// Consume messages from Kafka
	for message := range partitionConsumer.Messages() {
		messageID := string(message.Key)
		content := string(message.Value)
		processMessage(messageID, content)
	}
}
