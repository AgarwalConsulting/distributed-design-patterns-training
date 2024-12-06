package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work with a priority
type Task struct {
	id       int
	priority int // Higher values indicate higher priority
	action   func()
	index    int // Index in the priority queue
}

// PriorityQueue implements a priority queue for tasks
type PriorityQueue []*Task

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// Higher priority comes first
	return pq[i].priority > pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push adds a task to the queue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	task := x.(*Task)
	task.index = n
	*pq = append(*pq, task)
}

// Pop removes and returns the highest-priority task
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	task := old[n-1]
	old[n-1] = nil // Avoid memory leak
	task.index = -1
	*pq = old[0 : n-1]
	return task
}

// Scheduler manages the execution of tasks
type Scheduler struct {
	taskQueue PriorityQueue
	mu        sync.Mutex
	cond      *sync.Cond
	stop      chan struct{}
}

// NewScheduler creates a new Scheduler
func NewScheduler() *Scheduler {
	s := &Scheduler{
		taskQueue: make(PriorityQueue, 0),
		stop:      make(chan struct{}),
	}
	s.cond = sync.NewCond(&s.mu)
	heap.Init(&s.taskQueue)
	return s
}

// AddTask adds a new task to the scheduler
func (s *Scheduler) AddTask(task *Task) {
	s.mu.Lock()
	defer s.mu.Unlock()
	heap.Push(&s.taskQueue, task)
	s.cond.Signal() // Notify that a new task is available
}

// UpdatePriority dynamically updates the priority of a task
func (s *Scheduler) UpdatePriority(task *Task, newPriority int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	task.priority = newPriority
	heap.Fix(&s.taskQueue, task.index) // Adjust the position of the task in the queue
}

// Run starts the scheduler
func (s *Scheduler) Run() {
	for {
		s.mu.Lock()
		for len(s.taskQueue) == 0 {
			s.cond.Wait() // Wait until a task is available
		}
		task := heap.Pop(&s.taskQueue).(*Task)
		s.mu.Unlock()

		// Execute the task
		task.action()
	}
}

// Stop gracefully stops the scheduler
func (s *Scheduler) Stop() {
	close(s.stop)
}

func main() {
	scheduler := NewScheduler()

	// Start the scheduler
	go scheduler.Run()

	// Create and add tasks with different priorities
	scheduler.AddTask(&Task{
		id:       1,
		priority: 1,
		action: func() {
			fmt.Println("Executing Task 1 (Low Priority)")
		},
	})

	scheduler.AddTask(&Task{
		id:       2,
		priority: 5,
		action: func() {
			fmt.Println("Executing Task 2 (High Priority)")
		},
	})

	// Add a task that gets its priority updated dynamically
	task3 := &Task{
		id:       3,
		priority: 2,
		action: func() {
			fmt.Println("Executing Task 3 (Updated Priority)")
		},
	}
	scheduler.AddTask(task3)

	// Dynamically update the priority of Task 3
	time.Sleep(500 * time.Millisecond) // Simulate a delay
	fmt.Println("Updating Task 3 priority...")
	scheduler.UpdatePriority(task3, 10)

	// Give the scheduler time to process tasks
	time.Sleep(2 * time.Second)
	scheduler.Stop()
}
