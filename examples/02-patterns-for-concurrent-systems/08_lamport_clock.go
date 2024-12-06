package main

import "fmt"

type LamportClock struct {
	time int
}

func (lc *LamportClock) Increment() {
	lc.time++
}

func (lc *LamportClock) Update(other int) {
	if other > lc.time {
		lc.time = other
	}
	lc.time++
}

func (lc *LamportClock) Time() int {
	return lc.time
}

func main() {
	lc := &LamportClock{}
	lc.Increment()
	fmt.Println("Time:", lc.Time())
	lc.Update(5)
	fmt.Println("Time:", lc.Time())
}
