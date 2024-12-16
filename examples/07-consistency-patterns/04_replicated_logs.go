package main

import (
	"fmt"
)

type Node struct {
	nodeID string
	log    []string
}

func (n *Node) appendLog(entry string) int {
	n.log = append(n.log, entry)
	return len(n.log) - 1 // Return the index of the log entry
}

type Leader struct {
	Node
	followers []*Follower
}

func (l *Leader) replicateLog(entry string) int {
	index := l.appendLog(entry)
	for _, follower := range l.followers {
		follower.receiveLog(entry, index)
	}
	return index
}

type Follower struct {
	Node
}

func (f *Follower) receiveLog(entry string, index int) {
	if index == len(f.log) { // Ensure log order consistency
		f.appendLog(entry)
		fmt.Printf("Follower %s replicated entry: %s\n", f.nodeID, entry)
	}
}

func main() {
	follower1 := &Follower{Node: Node{nodeID: "Follower1"}}
	follower2 := &Follower{Node: Node{nodeID: "Follower2"}}
	leader := &Leader{
		Node:      Node{nodeID: "Leader"},
		followers: []*Follower{follower1, follower2},
	}

	leader.replicateLog("Log Entry 1")
	leader.replicateLog("Log Entry 2")
}
