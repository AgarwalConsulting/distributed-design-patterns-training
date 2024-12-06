package main

import (
	"fmt"
	"sync"
)

type Monitor struct {
	mu sync.Mutex
}

func (m *Monitor) SafeAction(action func()) {
	m.mu.Lock()
	defer m.mu.Unlock()
	action()
}

func main() {
	monitor := &Monitor{}

	monitor.SafeAction(func() {
		fmt.Println("Executing critical section")
	})
}
