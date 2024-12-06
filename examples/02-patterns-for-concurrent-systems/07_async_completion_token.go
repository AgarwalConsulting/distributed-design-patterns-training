package main

import (
	"fmt"
	"time"
)

func asyncTask() <-chan int {
	result := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		result <- 42
		close(result)
	}()
	return result
}

func main() {
	token := asyncTask()
	fmt.Println("Doing other work...")
	fmt.Println("Result:", <-token)
}
