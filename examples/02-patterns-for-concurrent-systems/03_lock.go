package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var counter int

	mu.Lock()
	counter++
	fmt.Println("Counter:", counter)
	mu.Unlock()
}
