package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	id            int
	state         string // "Follower", "Leader", "Candidate"
	term          int
	log           []string
	votedFor      int
	voteCount     int
	heartbeatTime time.Duration
	electionTimer *time.Timer
}

var peers []Node

func (n *Node) startElection() {
	n.state = "Candidate"
	n.term++
	n.voteCount = 1 // Vote for itself
	fmt.Printf("Node %d starts election with term %d\n", n.id, n.term)
	for _, peer := range peers {
		if peer.requestVote(n.term) {
			n.voteCount++
		}
	}
	if n.voteCount > len(peers)/2 {
		n.state = "Leader"
		fmt.Printf("Node %d became the leader for term %d\n", n.id, n.term)
		n.sendHeartbeats()
	} else {
		n.state = "Follower"
	}
}

func (n *Node) requestVote(term int) bool {
	if term > n.term {
		n.term = term
		n.state = "Follower"
		n.votedFor = n.id
		return true
	}
	return false
}

func (n *Node) sendHeartbeats() {
	fmt.Printf("Leader %d sending heartbeats\n", n.id)
	for _, peer := range peers {
		peer.receiveHeartbeat(n.term)
	}
}

func (n *Node) receiveHeartbeat(term int) {
	if term >= n.term {
		n.term = term
		n.state = "Follower"
		fmt.Printf("Node %d received heartbeat from Leader %d\n", n.id, n.term)
	}
}

func (n *Node) onTimeout() {
	if n.state == "Follower" {
		n.startElection()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// Simulate nodes
	peers = []Node{
		{id: 1, state: "Follower", term: 0},
		{id: 2, state: "Follower", term: 0},
		{id: 3, state: "Follower", term: 0},
	}

	// Simulate timeouts and leader election
	for _, node := range peers {
		go func(n Node) {
			electionTimeout := rand.Intn(150) + 150 // Random timeout between 150ms and 300ms
			time.Sleep(time.Duration(electionTimeout) * time.Millisecond)
			n.onTimeout()
		}(node)
	}

	// Run forever
	select {}
}
