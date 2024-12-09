package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Order and Outbox event structs
type Order struct {
	ID       uint    `json:"id"`
	Customer string  `json:"customer"`
	Total    float64 `json:"total"`
}

type OutboxEvent struct {
	ID        uint   `json:"id"`
	EventType string `json:"event_type"`
	OrderID   uint   `json:"order_id"`
}

func createOrderWithEvent(db *gorm.DB, order Order) {
	// Start a transaction
	tx := db.Begin()

	// Create order in the database
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		log.Fatal("Error creating order:", err)
	}

	// Create outbox event
	event := OutboxEvent{EventType: "ORDER_PLACED", OrderID: order.ID}
	if err := tx.Create(&event).Error; err != nil {
		tx.Rollback()
		log.Fatal("Error creating outbox event:", err)
	}

	// Commit the transaction
	tx.Commit()
	fmt.Println("Order created and outbox event saved:", order)
}

func main() {
	// Set up SQLite database for both orders and outbox
	db, err := gorm.Open("sqlite3", "./order_outbox.db")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	db.AutoMigrate(&Order{}, &OutboxEvent{}) // Migrate models

	// Create order and event in a transactional manner
	order := Order{Customer: "Alice", Total: 150.00}
	createOrderWithEvent(db, order)
}
