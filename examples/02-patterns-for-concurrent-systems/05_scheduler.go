package main

import (
	"fmt"
	"time"
)

func scheduler(tasks []func()) {
	for {
		for _, task := range tasks {
			task()
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	tasks := []func(){
		func() { fmt.Println("Task 1") },
		func() { fmt.Println("Task 2") },
		func() { fmt.Println("Task 3") },
	}
	go scheduler(tasks)
	time.Sleep(5 * time.Second)
}
