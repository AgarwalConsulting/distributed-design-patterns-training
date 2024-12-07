package main

import (
	"fmt"
	"log"
)

type Step func() error

func OrderBooking() error {
	// Simulate order booking
	fmt.Println("Booking order...")
	return nil
}

func Payment() error {
	// Simulate payment
	fmt.Println("Processing payment...")
	return nil
}

func Shipping() error {
	// Simulate shipping
	fmt.Println("Shipping order...")
	return nil
}

func CancelShipping() error {
	// Simulate cancel shipping
	fmt.Println("Canceling shipping...")
	return nil
}

func CancelPayment() error {
	// Simulate cancel payment
	fmt.Println("Cancelling payment...")
	return nil
}

func RunSaga(steps []Step, compensations []Step) error {
	for i, step := range steps {
		err := step()
		if err != nil {
			fmt.Println("Error occurred, starting compensations...")
			// Execute compensating actions in reverse order
			for j := i - 1; j >= 0; j-- {
				compensations[j]()
			}
			return err
		}
	}
	return nil
}

func main() {
	steps := []Step{OrderBooking, Payment, Shipping}
	compensations := []Step{CancelShipping, CancelPayment}

	err := RunSaga(steps, compensations)
	if err != nil {
		log.Println("Saga failed:", err)
	} else {
		fmt.Println("Saga completed successfully.")
	}
}
