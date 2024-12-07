package main

import (
	"fmt"
)

type Event interface {
	ApplyTo(*BankAccount)
}

type Deposit struct {
	Amount float64
}

func (d Deposit) ApplyTo(account *BankAccount) {
	account.Balance += d.Amount
}

type Withdraw struct {
	Amount float64
}

func (w Withdraw) ApplyTo(account *BankAccount) {
	account.Balance -= w.Amount
}

type BankAccount struct {
	Balance float64
	Events  []Event
}

func (account *BankAccount) ApplyEvent(event Event) {
	event.ApplyTo(account)
	account.Events = append(account.Events, event)
}

func NewBankAccount() *BankAccount {
	return &BankAccount{}
}

func main() {
	account := NewBankAccount()

	// Apply events
	account.ApplyEvent(Deposit{Amount: 100.0})
	account.ApplyEvent(Withdraw{Amount: 50.0})

	// Rebuild the state from events
	reconstructedAccount := NewBankAccount()
	for _, event := range account.Events {
		event.ApplyTo(reconstructedAccount)
	}

	fmt.Printf("Reconstructed Balance: %.2f\n", reconstructedAccount.Balance)
}
