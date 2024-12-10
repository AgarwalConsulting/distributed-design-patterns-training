package main

import (
	"fmt"
	"hash/fnv"
)

// Hash-based sharding
func shard(userID string, numShards int) int {
	h := fnv.New32a()
	h.Write([]byte(userID))
	return int(h.Sum32()) % numShards
}

func main() {
	// Simulating 3 database shards
	shards := map[int][]string{
		0: {}, // Shard 0
		1: {}, // Shard 1
		2: {}, // Shard 2
	}

	users := []string{"user1", "user2", "user3", "user4", "user5"}
	for _, user := range users {
		shardID := shard(user, len(shards))
		shards[shardID] = append(shards[shardID], user)
	}

	// Display users in each shard
	for id, users := range shards {
		fmt.Printf("Shard %d: %v\n", id, users)
	}
}
