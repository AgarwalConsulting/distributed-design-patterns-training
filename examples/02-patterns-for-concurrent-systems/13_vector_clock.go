package main

import (
	"fmt"
	"sync"
)

type VectorClock struct {
	clock map[string]int
	lock  sync.Mutex
}

// NewVectorClock initializes a new vector clock for a given process.
func NewVectorClock() *VectorClock {
	return &VectorClock{
		clock: make(map[string]int),
	}
}

// Increment increments the vector clock for the current process.
func (vc *VectorClock) Increment(processID string) {
	vc.lock.Lock()
	defer vc.lock.Unlock()
	vc.clock[processID]++
}

// Update merges another vector clock into the current one.
func (vc *VectorClock) Update(other *VectorClock) {
	vc.lock.Lock()
	defer vc.lock.Unlock()

	for processID, timestamp := range other.clock {
		if vc.clock[processID] < timestamp {
			vc.clock[processID] = timestamp
		}
	}
}

// SendMessage simulates sending a message by attaching the current vector clock.
func (vc *VectorClock) SendMessage() map[string]int {
	vc.lock.Lock()
	defer vc.lock.Unlock()
	copy := make(map[string]int)
	for k, v := range vc.clock {
		copy[k] = v
	}
	return copy
}

// ReceiveMessage simulates receiving a message and updating the vector clock.
func (vc *VectorClock) ReceiveMessage(message map[string]int, processID string) {
	vc.lock.Lock()
	defer vc.lock.Unlock()

	for processID, timestamp := range message {
		if vc.clock[processID] < timestamp {
			vc.clock[processID] = timestamp
		}
	}
	vc.clock[processID]++
}

// Display shows the current state of the vector clock.
func (vc *VectorClock) Display() {
	vc.lock.Lock()
	defer vc.lock.Unlock()
	fmt.Println(vc.clock)
}

func main() {
	// Create vector clocks for two processes.
	vc1 := NewVectorClock()
	vc2 := NewVectorClock()

	// Process IDs
	p1 := "P1"
	p2 := "P2"

	// Simulate events.
	vc1.Increment(p1) // P1 performs an event.
	vc1.Display()

	message := vc1.SendMessage()    // P1 sends a message to P2.
	vc2.ReceiveMessage(message, p2) // P2 receives the message.
	vc2.Display()

	vc2.Increment(p2) // P2 performs another event.
	vc2.Display()

	message = vc2.SendMessage()     // P2 sends a message to P1.
	vc1.ReceiveMessage(message, p1) // P1 receives the message.
	vc1.Display()
}
