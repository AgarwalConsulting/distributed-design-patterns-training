package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		mutex  sync.Mutex
		cond   = sync.NewCond(&mutex)
		buffer []int
	)

	const bufferSize = 5

	// Consumer goroutine
	go func() {
		for {
			mutex.Lock()
			for len(buffer) == 0 { // Wait while buffer is empty
				cond.Wait()
			}

			// Consume an item
			item := buffer[0]
			buffer = buffer[1:]
			fmt.Println("Consumed:", item)

			mutex.Unlock()

			// Simulate processing time
			time.Sleep(time.Second)
		}
	}()

	// Producer goroutine
	go func() {
		for i := 1; i <= 10; i++ {
			mutex.Lock()
			for len(buffer) == bufferSize { // Wait while buffer is full
				cond.Wait()
			}

			// Produce an item
			buffer = append(buffer, i)
			fmt.Println("Produced:", i)

			cond.Signal() // Signal that a new item is available
			mutex.Unlock()

			// Simulate production time
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Allow the program to run for a while
	time.Sleep(15 * time.Second)
}
