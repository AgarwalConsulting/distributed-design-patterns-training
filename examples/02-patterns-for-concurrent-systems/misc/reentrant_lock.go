package main

import (
	"fmt"
	"runtime"
	"sync"
)

// ReentrantLock is a custom reentrant lock implementation
type ReentrantLock struct {
	mu         sync.Mutex
	owner      int64
	recursions int
}

// Lock acquires the lock, or increases recursion count if already owned
func (rl *ReentrantLock) Lock() {
	gid := getGID() // Get the current goroutine ID

	rl.mu.Lock()
	defer rl.mu.Unlock()

	if rl.owner == gid {
		rl.recursions++
		return
	}

	// Wait until the lock is free
	for rl.owner != 0 {
		rl.mu.Unlock()
		rl.mu.Lock()
	}

	rl.owner = gid
	rl.recursions = 1
}

// Unlock releases the lock, or decreases recursion count if owned multiple times
func (rl *ReentrantLock) Unlock() {
	gid := getGID()

	rl.mu.Lock()
	defer rl.mu.Unlock()

	if rl.owner != gid {
		panic("unlock called by non-owner")
	}

	rl.recursions--
	if rl.recursions == 0 {
		rl.owner = 0
	}
}

// getGID retrieves the ID of the current goroutine
func getGID() int64 {
	// This function uses runtime calls to fetch the goroutine ID.
	// It's not part of the Go standard library and is for demonstration purposes.
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	var gid int64
	fmt.Sscanf(string(buf[:n]), "goroutine %d ", &gid)
	return gid
}

func main() {
	var rl ReentrantLock

	rl.Lock()
	fmt.Println("Acquired lock the first time")

	rl.Lock()
	fmt.Println("Acquired lock the second time (reentrant)")

	rl.Unlock()
	fmt.Println("Released lock once")

	rl.Unlock()
	fmt.Println("Released lock completely")
}
