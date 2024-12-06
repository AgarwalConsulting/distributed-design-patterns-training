package main

import (
	"fmt"
	"time"
)

type Reactor struct {
	events chan string
}

func NewReactor() *Reactor {
	return &Reactor{events: make(chan string)}
}

func (r *Reactor) Handle(event string, handler func()) {
	go func() {
		for e := range r.events {
			if e == event {
				handler()
			}
		}
	}()
}

func (r *Reactor) Emit(event string) {
	r.events <- event
}

func main() {
	reactor := NewReactor()
	reactor.Handle("ping", func() {
		fmt.Println("Pong!")
	})

	go func() {
		time.Sleep(1 * time.Second)
		reactor.Emit("ping")
	}()
	time.Sleep(2 * time.Second)
}
