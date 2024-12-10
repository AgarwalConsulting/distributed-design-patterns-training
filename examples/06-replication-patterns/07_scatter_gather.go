package main

import (
	"fmt"
	"sync"
)

func search(node string, query string, results chan<- string) {
	// Simulate search processing
	results <- fmt.Sprintf("Result from %s for query '%s'", node, query)
}

func main() {
	nodes := []string{"node1", "node2", "node3"}
	query := "example query"
	results := make(chan string, len(nodes))
	var wg sync.WaitGroup

	// Scatter
	for _, node := range nodes {
		wg.Add(1)
		go func(n string) {
			defer wg.Done()
			search(n, query, results)
		}(node)
	}

	// Wait for all nodes
	go func() {
		wg.Wait()
		close(results)
	}()

	// Gather
	fmt.Println("Search Results:")
	for res := range results {
		fmt.Println(res)
	}
}
