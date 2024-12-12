package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := &Counter{}
	var wg sync.WaitGroup

	// Simulating concurrent increments on the counter
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	// Wait for all increments to complete
	wg.Wait()

	// Simulate a read operation
	fmt.Println("Counter value after 5 increments:", counter.GetValue())
}
