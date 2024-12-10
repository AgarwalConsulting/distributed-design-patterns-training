package main

import (
	"fmt"
	"time"
)

var leader string

func leaderElection(nodes []string) {
	for {
		leader = nodes[time.Now().UnixNano()%int64(len(nodes))]
		fmt.Printf("New leader elected: %s\n", leader)
		time.Sleep(5 * time.Second)
	}
}

func main() {
	nodes := []string{"node1", "node2", "node3"}
	go leaderElection(nodes)

	select {}
}
