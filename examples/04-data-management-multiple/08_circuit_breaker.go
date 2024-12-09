package main

import (
	"fmt"
	"log"

	"github.com/sony/gobreaker"
)

// Function to simulate a service call
func callService() (string, error) {
	// Simulate failure
	return "", fmt.Errorf("service is down")
}

func main() {
	// Configure circuit breaker
	settings := gobreaker.Settings{
		Name:    "Service Circuit Breaker",
		Timeout: 5,
	}
	circuit := gobreaker.NewCircuitBreaker(settings)

	// Making a service call using the circuit breaker
	result, err := circuit.Execute(func() (interface{}, error) {

		return callService()
	})

	if err != nil {
		log.Println("Circuit breaker triggered:", err)
		return
	}

	// Process the result if no error
	fmt.Println("Service call result:", result)
}
