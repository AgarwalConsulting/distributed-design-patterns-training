package main

import (
	"fmt"
)

type Coordinator struct {
	participants []Participant
}

type Participant interface {
	Prepare() bool
	Commit() bool
	Rollback() bool
}

type DatabaseA struct{}
type DatabaseB struct{}

func (d DatabaseA) Prepare() bool {
	fmt.Println("Database A is preparing...")
	return true // simulate success
}

func (d DatabaseA) Commit() bool {
	fmt.Println("Database A commits.")
	return true
}

func (d DatabaseA) Rollback() bool {
	fmt.Println("Database A rolls back.")
	return true
}

func (d DatabaseB) Prepare() bool {
	fmt.Println("Database B is preparing...")
	return true // simulate success
}

func (d DatabaseB) Commit() bool {
	fmt.Println("Database B commits.")
	return true
}

func (d DatabaseB) Rollback() bool {
	fmt.Println("Database B rolls back.")
	return true
}

func (c *Coordinator) StartTransaction() {
	// Phase 1: Prepare
	allReady := true
	for _, participant := range c.participants {
		if !participant.Prepare() {
			allReady = false
			break
		}
	}

	// Phase 2: Commit or Rollback
	if allReady {
		fmt.Println("All participants are ready, committing.")
		for _, participant := range c.participants {
			if !participant.Commit() {
				participant.Rollback()
			}
		}
	} else {
		fmt.Println("Participants are not ready, rolling back.")
		for _, participant := range c.participants {
			participant.Rollback()
		}
	}
}

func main() {
	// Create participants
	dbA := DatabaseA{}
	dbB := DatabaseB{}

	// Create coordinator and start transaction
	coordinator := &Coordinator{participants: []Participant{dbA, dbB}}
	coordinator.StartTransaction()
}
