package main

import (
	"fmt"
)

// DomainEvent is a simple interface for events.
type DomainEvent interface {
	Name() string
}

// UserRegisteredEvent represents a user registration event.
type UserRegisteredEvent struct {
	UserID string
	Email  string
}

func (e UserRegisteredEvent) Name() string {
	return "UserRegisteredEvent"
}

// EventHandler handles domain events.
type EventHandler func(event DomainEvent)

// EventBus allows publishing and subscribing to domain events.
type EventBus struct {
	handlers map[string][]EventHandler
}

func NewEventBus() *EventBus {
	return &EventBus{handlers: make(map[string][]EventHandler)}
}

func (bus *EventBus) Subscribe(eventType string, handler EventHandler) {
	bus.handlers[eventType] = append(bus.handlers[eventType], handler)
}

func (bus *EventBus) Publish(event DomainEvent) {
	for _, handler := range bus.handlers[event.Name()] {
		handler(event)
	}
}

func main() {
	bus := NewEventBus()

	// Subscribe to events
	bus.Subscribe("UserRegisteredEvent", func(event DomainEvent) {
		e := event.(UserRegisteredEvent)
		fmt.Printf("Welcome email sent to %s\n", e.Email)
	})

	// Publish an event
	event := UserRegisteredEvent{UserID: "123", Email: "user@example.com"}
	bus.Publish(event)
}
