package main

import "fmt"

type ActiveObject struct {
	requests chan func()
}

func NewActiveObject() *ActiveObject {
	ao := &ActiveObject{
		requests: make(chan func()),
	}
	go ao.run()
	return ao
}

func (ao *ActiveObject) run() {
	for req := range ao.requests {
		req()
	}
}

func (ao *ActiveObject) Do(action func()) {
	ao.requests <- action
}

func main() {
	ao := NewActiveObject()
	ao.Do(func() {
		fmt.Println("Task 1")
	})
	ao.Do(func() {
		fmt.Println("Task 2")
	})
}
