package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(time.Second) // Simulate work
		results <- task * 2
	}
}

func main() {
	tasks := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup

	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, tasks, results)
		}(i)
	}

	// Add tasks to the queue
	go func() {
		for i := 1; i <= 5; i++ {
			tasks <- i
		}
		close(tasks)
	}()

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}
