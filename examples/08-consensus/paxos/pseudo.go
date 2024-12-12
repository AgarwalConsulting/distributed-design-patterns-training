package main

import "fmt"

type Proposer struct {
	proposalNumber int
	value          string
	acceptors      []Acceptor
}

func (p *Proposer) propose() {
	responses := []string{}
	for _, acceptor := range p.acceptors {
		response := acceptor.prepare(p.proposalNumber)
		if response != "" {
			responses = append(responses, response)
		}
	}

	if len(responses) > len(p.acceptors)/2 {
		highestValue := p.value
		if len(responses) > 0 {
			highestValue = responses[0] // Simplified to take the first response
		}
		for _, acceptor := range p.acceptors {
			acceptor.accept(p.proposalNumber, highestValue)
		}
	}
}

type Acceptor struct {
	highestProposalNumber int
	acceptedValue         string
}

func (a *Acceptor) prepare(proposalNumber int) string {
	if proposalNumber > a.highestProposalNumber {
		a.highestProposalNumber = proposalNumber
		return a.acceptedValue
	}
	return ""
}

func (a *Acceptor) accept(proposalNumber int, value string) {
	if proposalNumber >= a.highestProposalNumber {
		a.acceptedValue = value
		a.highestProposalNumber = proposalNumber
	}
}

func main() {
	acceptors := []Acceptor{{}, {}, {}}
	proposer := Proposer{proposalNumber: 1, value: "value1", acceptors: acceptors}
	proposer.propose()

	fmt.Println("Accepted value:", acceptors[0].acceptedValue)
}
