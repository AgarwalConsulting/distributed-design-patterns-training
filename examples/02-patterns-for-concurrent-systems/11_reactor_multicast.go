package main

import (
	"fmt"
	"sync"
)

type Event struct {
	Data string
}

type Reactor struct {
	handlerEvents map[string]chan Event
	mu            sync.Mutex
}

func NewReactor() *Reactor {
	return &Reactor{
		handlerEvents: make(map[string]chan Event),
	}
}

func (r *Reactor) RegisterHandler(name string) chan Event {
	r.mu.Lock()
	defer r.mu.Unlock()
	ch := make(chan Event, 10) // Buffered channel for each handler
	r.handlerEvents[name] = ch
	return ch
}

func (r *Reactor) Dispatch(event Event) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for name, ch := range r.handlerEvents {
		select {
		case ch <- event: // Non-blocking send
			fmt.Printf("Event sent to handler %s\n", name)
		default:
			fmt.Printf("Handler %s is too slow, dropping event\n", name)
		}
	}
}
