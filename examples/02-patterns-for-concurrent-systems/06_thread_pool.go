package main

import (
	"fmt"
	"sync"
)

type Task func()

type ThreadPool struct {
	tasks chan Task
	wg    sync.WaitGroup
}

func NewThreadPool(size int) *ThreadPool {
	pool := &ThreadPool{tasks: make(chan Task)}
	for i := 0; i < size; i++ {
		go pool.worker()
	}
	return pool
}

func (p *ThreadPool) worker() {
	for task := range p.tasks {
		task()
		p.wg.Done()
	}
}

func (p *ThreadPool) Submit(task Task) {
	p.wg.Add(1)
	p.tasks <- task
}

func (p *ThreadPool) Wait() {
	p.wg.Wait()
	close(p.tasks)
}

func main() {
	pool := NewThreadPool(3)
	for i := 0; i < 10; i++ {
		taskNum := i
		pool.Submit(func() {
			fmt.Println("Executing task", taskNum)
		})
	}
	pool.Wait()
}
