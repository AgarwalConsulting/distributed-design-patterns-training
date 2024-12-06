package main

import (
	"fmt"
	"time"
)

type HybridClock struct {
	logical int
}

func (hc *HybridClock) Now() string {
	hc.logical++
	return fmt.Sprintf("%d|%d", time.Now().UnixNano(), hc.logical)
}

func main() {
	hc := &HybridClock{}
	fmt.Println("Time:", hc.Now())
}
