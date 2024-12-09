package main

import (
	"fmt"
	"sync"
)

// Step 1: Validate order
func validateOrder(orderID int) string {
	return fmt.Sprintf("Order %d validated", orderID)
}

// Step 2: Process payment
func processPayment(orderID int) string {
	return fmt.Sprintf("Payment for order %d processed", orderID)
}

// Step 3: Update inventory
func updateInventory(orderID int) string {
	return fmt.Sprintf("Inventory updated for order %d", orderID)
}

// Orchestrator for handling the order
func orchestrateOrder(orderID int) {
	var wg sync.WaitGroup
	results := make(chan string, 3)

	// Step 1: Validate Order
	wg.Add(1)
	go func() {
		defer wg.Done()
		results <- validateOrder(orderID)
	}()

	// Step 2: Process Payment
	wg.Add(1)
	go func() {
		defer wg.Done()
		results <- processPayment(orderID)
	}()

	// Step 3: Update Inventory
	wg.Add(1)
	go func() {
		defer wg.Done()
		results <- updateInventory(orderID)
	}()

	// Wait for all steps to complete
	wg.Wait()
	close(results)

	// Print results
	fmt.Println("Order Orchestration Results:")
	for result := range results {
		fmt.Println(result)
	}
}

func main() {
	fmt.Println("Starting Orchestration Example: E-Commerce Order")
	orchestrateOrder(101)
}
