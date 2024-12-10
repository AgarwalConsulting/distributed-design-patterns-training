package main

import (
	"fmt"
	"sync"
	"time"
)

var leaderData = "Leader data"
var followerData = "Follower data (lagged)"
var mu sync.RWMutex

func writeToLeader(data string) {
	mu.Lock()
	defer mu.Unlock()
	leaderData = data
	followerData = data                // Simulate delayed replication
	time.Sleep(100 * time.Millisecond) // Simulate lag
}

func readFromFollower() string {
	mu.RLock()
	defer mu.RUnlock()
	return followerData
}

func main() {
	go func() {
		writeToLeader("New data from leader")
	}()

	time.Sleep(200 * time.Millisecond) // Wait for replication
	fmt.Println("Reading from follower:", readFromFollower())
}
