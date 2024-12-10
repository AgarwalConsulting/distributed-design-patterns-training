// Install: go get github.com/hashicorp/raft
package main

import (
	"log"
	"os"
	"time"

	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
)

func main() {
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID("node1")

	transport, err := raft.NewTCPTransport("127.0.0.1:8080", nil, 3, time.Second, os.Stdout)
	if err != nil {
		log.Fatalf("Failed to create transport: %v", err)
	}

	logStore, err := raftboltdb.NewBoltStore("raft-log.db")
	if err != nil {
		log.Fatalf("Failed to create log store: %v", err)
	}

	stableStore, err := raftboltdb.NewBoltStore("raft-stable.db")
	if err != nil {
		log.Fatalf("Failed to create stable store: %v", err)
	}

	snapshotStore, err := raft.NewFileSnapshotStore(".", 2, os.Stdout)
	if err != nil {
		log.Fatalf("Failed to create snapshot store: %v", err)
	}

	r, err := raft.NewRaft(config, nil, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		log.Fatalf("Failed to create Raft instance: %v", err)
	}

	log.Println("Raft instance running. Waiting for leader election...")
	select {} // Keep the process running
}
