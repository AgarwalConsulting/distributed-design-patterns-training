package main

import (
	"fmt"
	"time"
)

// Event channels
var orderValidated = make(chan int)
var paymentProcessed = make(chan int)

// Step 1: Validate order
func validateOrder(orderID int) {
	fmt.Printf("Validating order %d...\n", orderID)
	time.Sleep(1 * time.Second) // Simulate processing time
	fmt.Printf("Order %d validated.\n", orderID)
	orderValidated <- orderID // Trigger next step
}

// Step 2: Process payment
func processPayment() {
	for orderID := range orderValidated {
		fmt.Printf("Processing payment for order %d...\n", orderID)
		time.Sleep(1 * time.Second) // Simulate processing time
		fmt.Printf("Payment processed for order %d.\n", orderID)
		paymentProcessed <- orderID // Trigger next step
	}
}

// Step 3: Update inventory
func updateInventory() {
	for orderID := range paymentProcessed {
		fmt.Printf("Updating inventory for order %d...\n", orderID)
		time.Sleep(1 * time.Second) // Simulate processing time
		fmt.Printf("Inventory updated for order %d.\n", orderID)
	}
}

func main() {
	fmt.Println("Starting Choreography Example: E-Commerce Order")

	// Start services
	go processPayment()
	go updateInventory()

	// Simulate an order being placed
	go validateOrder(101)

	// Prevent program from exiting immediately
	time.Sleep(5 * time.Second)
}
