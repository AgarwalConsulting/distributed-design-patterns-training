package main

import (
	"fmt"
)

type Proctor struct {
	checks []func() bool
}

func (p *Proctor) AddCheck(check func() bool) {
	p.checks = append(p.checks, check)
}

func (p *Proctor) RunChecks() {
	for i, check := range p.checks {
		if check() {
			fmt.Printf("Check %d passed\n", i)
		} else {
			fmt.Printf("Check %d failed\n", i)
		}
	}
}

func main() {
	proctor := &Proctor{}
	proctor.AddCheck(func() bool { return true })
	proctor.AddCheck(func() bool { return false })
	proctor.RunChecks()
}
