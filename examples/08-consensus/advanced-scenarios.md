# Advanced Go implementations for **Paxos** and **Raft**

1. **Leader Election with Network Partitioning**
2. **Log Compaction and Snapshotting**
3. **Handling Node Failures Gracefully**
4. **Advanced Failure Recovery in Raft**

---

## 1. Advanced Leader Election with Network Partitioning

Network partitioning is a common issue in distributed systems, and how **Raft** and **Paxos** handle it can greatly affect system availability and consistency.

### Raft Network Partitioning Handling

In **Raft**, the **leader** and **majority** rule ensure that a partition where a majority of nodes are unreachable cannot make progress. The new **election** mechanism after partitioning is key.

#### Raft Leader Election with Partitioning Example

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	id         int
	state      string // "Follower", "Leader", "Candidate"
	term       int
	voteCount  int
	votedFor   int
	log        []string
	electionTimer *time.Timer
}

var peers []Node

func (n *Node) startElection() {
	n.state = "Candidate"
	n.term++
	n.voteCount = 1 // Vote for itself
	fmt.Printf("Node %d starts election with term %d\n", n.id, n.term)

	// Request votes from other nodes
	for _, peer := range peers {
		if peer.requestVote(n.term) {
			n.voteCount++
		}
	}

	// Check if the candidate has a majority
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
	// Send heartbeats to followers
	for _, peer := range peers {
		if peer.state != "Leader" {
			peer.receiveHeartbeat(n.term)
		}
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

	// Simulate network partition and timeouts
	for _, node := range peers {
		go func(n Node) {
			electionTimeout := rand.Intn(150) + 150 // Random timeout between 150ms and 300ms
			time.Sleep(time.Duration(electionTimeout) * time.Millisecond)
			n.onTimeout()
		}(node)
	}

	// Simulate nodes healing the partition after some time
	time.Sleep(5 * time.Second)
	// Trigger network healing by checking state transitions
	for _, node := range peers {
		if node.state == "Leader" {
			fmt.Printf("Leader %d has been elected.\n", node.id)
		}
	}

	// Run forever
	select {}
}
```

### Key Notes*

- **Election Timeout**: Each node waits for a random timeout before starting the election.
- **Network Partitioning**: After partitioning, only the majority partition can elect a new leader.
- **Leader Heartbeats**: Heartbeats are sent by the leader to ensure followers stay up to date.

In the case of network partitions, the system ensures that only one leader is elected once the partition heals, and the majority partition controls the system. After healing, the nodes elect a new leader if necessary.

---

## 2. Log Compaction and Snapshotting in Raft

To prevent logs from growing indefinitely, **Raft** systems use **log compaction** via **snapshotting**. This process involves taking periodic snapshots of the state machine and discarding old log entries.

### Raft Snapshotting Example

```go
package main

import (
	"fmt"
	"time"
)

type Node struct {
	id          int
	state       string
	log         []string
	lastApplied int
}

var peers []Node

// Take snapshot to compact logs
func (n *Node) takeSnapshot() {
	n.lastApplied = len(n.log) // Track the last applied entry
	fmt.Printf("Node %d taking snapshot at log index %d\n", n.id, n.lastApplied)
	n.log = n.log[n.lastApplied:] // Compact log after snapshot
}

func (n *Node) appendLogEntry(entry string) {
	n.log = append(n.log, entry)
}

func main() {
	// Simulate Raft Nodes with log entries
	peers = []Node{
		{id: 1, state: "Follower", log: []string{}, lastApplied: 0},
		{id: 2, state: "Leader", log: []string{}, lastApplied: 0},
		{id: 3, state: "Follower", log: []string{}, lastApplied: 0},
	}

	// Append some log entries
	peers[1].appendLogEntry("Log entry 1")
	peers[1].appendLogEntry("Log entry 2")

	// Take snapshot periodically
	peers[1].takeSnapshot()

	// Add more log entries after snapshot
	peers[1].appendLogEntry("Log entry 3")

	// Print log after snapshot
	fmt.Printf("Logs after snapshot: %v\n", peers[1].log)
}
```

### Key Notes

- **Snapshotting**: After taking a snapshot, Raft nodes keep track of the last applied index, and logs older than the snapshot are discarded.
- **Log Size Reduction**: Snapshotting reduces the storage footprint of logs, ensuring the system doesn’t run into memory issues with growing logs.

---

## 3. Handling Node Failures Gracefully in Raft

In a **Raft**-based system, nodes can fail and restart without causing data loss, thanks to **log replication** and **leader election**. However, handling failures gracefully is crucial to maintaining consistency and availability.

### Handling Node Failures (Go Example)

```go
package main

import (
	"fmt"
	"time"
)

type Node struct {
	id          int
	state       string // "Follower", "Leader", "Candidate"
	term        int
	log         []string
	isFailed    bool
}

func (n *Node) startElection() {
	if n.isFailed {
		fmt.Printf("Node %d is failed and cannot start election\n", n.id)
		return
	}
	n.state = "Candidate"
	n.term++
	fmt.Printf("Node %d starts election for term %d\n", n.id, n.term)
	// Send election request to peers (simplified)
}

func (n *Node) receiveLogEntries(entries []string) {
	if n.isFailed {
		fmt.Printf("Node %d is failed and cannot receive log entries\n", n.id)
		return
	}
	n.log = append(n.log, entries...)
	fmt.Printf("Node %d received log entries: %v\n", n.id, entries)
}

func (n *Node) fail() {
	n.isFailed = true
	fmt.Printf("Node %d has failed\n", n.id)
}

func (n *Node) recover() {
	n.isFailed = false
	fmt.Printf("Node %d has recovered\n", n.id)
}

func main() {
	node := Node{id: 1, state: "Follower", term: 0, isFailed: false}

	// Simulating normal operation
	node.startElection()
	node.receiveLogEntries([]string{"entry1", "entry2"})

	// Simulate failure and recovery
	node.fail()
	node.startElection() // Shouldn't start election
	node.recover()
	node.startElection() // Should start election again

	// Simulating receiving logs after recovery
	node.receiveLogEntries([]string{"entry3", "entry4"})
}
```

### Key Notes

- **Failure Handling**: If a node is failed, it doesn't participate in the election or log replication.
- **Recovery**: Once a node recovers, it can rejoin the system, start new elections if needed, and synchronize its log with the leader.

---

## 4. Advanced Failure Recovery in Raft

Raft’s strength lies in its **leader election** and **log replication** mechanisms. If a node crashes and later recovers, it must synchronize with the leader's log to avoid inconsistencies.

### Recovery After Leader Failure (Advanced Example)

```go
package main

import (
	"fmt"
)

type Node struct {
	id         int
	state      string // "Leader", "Follower", "Candidate"
	term       int
	log        []string
	leaderId   int
	lastApplied int
}

func (n *Node

) crashAndRecover() {
	fmt.Printf("Node %d crashed\n", n.id)
	// Simulate crash (node losing state)
	n.state = "Follower"
	n.term = 0
	n.log = nil

	// Simulate recovery
	fmt.Printf("Node %d recovering\n", n.id)
	n.state = "Follower"
	n.term = n.leaderId // Sync term with leader
	n.log = []string{} // Synchronize logs
}

func main() {
	node := Node{id: 1, state: "Follower", term: 1, leaderId: 2}

	// Simulate log entries and leader information
	node.log = append(node.log, "Log entry 1", "Log entry 2")

	// Simulate node crash and recovery
	node.crashAndRecover()

	// Simulate post-recovery actions
	fmt.Printf("Node state: %s, Term: %d, Log: %v\n", node.state, node.term, node.log)
}
```

### Key Notes

- **Leader Syncing**: Upon recovery, the node syncs its term and log with the leader’s state.
- **No Data Loss**: Raft’s **log replication** ensures that the system’s state remains consistent across failures and recoveries.

---

### Conclusion

With **Paxos** and **Raft**, you can build robust distributed systems that handle failure gracefully. **Raft** particularly excels in ensuring **consistency** and **availability** in the face of node failures, partitions, and leader crashes, with sophisticated mechanisms for **leader election**, **log replication**, and **failure recovery**.
