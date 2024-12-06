package main

import "fmt"

type GenerationClock struct {
	generation int
}

func (gc *GenerationClock) NextGeneration() {
	gc.generation++
}

func (gc *GenerationClock) Current() int {
	return gc.generation
}

func main() {
	gc := &GenerationClock{}
	fmt.Println("Generation:", gc.Current())
	gc.NextGeneration()
	fmt.Println("Generation:", gc.Current())
}
