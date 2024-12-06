package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Service represents a microservice with health data.
type Service struct {
	Name   string
	Status string // "Healthy", "Unhealthy", or "Unknown"
}

// Proctor monitors and manages the health of services.
type Proctor struct {
	services map[string]*Service
	mu       sync.Mutex
	alerts   chan string
	stop     chan struct{}
	wg       sync.WaitGroup
}

// NewProctor creates a new Proctor.
func NewProctor() *Proctor {
	return &Proctor{
		services: make(map[string]*Service),
		alerts:   make(chan string, 10),
		stop:     make(chan struct{}),
	}
}

// RegisterService registers a new service with the Proctor.
func (p *Proctor) RegisterService(name string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.services[name] = &Service{
		Name:   name,
		Status: "Unknown",
	}
}

// UpdateHealth updates the health status of a service.
func (p *Proctor) UpdateHealth(name string, status string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if service, exists := p.services[name]; exists {
		service.Status = status
	}
}

// Monitor continuously checks the health of services and triggers alerts.
func (p *Proctor) Monitor() {
	p.wg.Add(1)
	defer p.wg.Done()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			p.checkHealth()
		case <-p.stop:
			return
		}
	}
}

// checkHealth performs health checks and sends alerts if issues are found.
func (p *Proctor) checkHealth() {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, service := range p.services {
		if service.Status == "Unhealthy" {
			alert := fmt.Sprintf("ALERT: Service %s is unhealthy!", service.Name)
			fmt.Println(alert)
			p.alerts <- alert
		} else if service.Status == "Unknown" {
			fmt.Printf("WARNING: Service %s health status is unknown.\n", service.Name)
		}
	}
}

// Stop gracefully stops the Proctor.
func (p *Proctor) Stop() {
	close(p.stop)
	p.wg.Wait()
	close(p.alerts)
}

// SimulateService simulates random health updates for a service.
func SimulateService(p *Proctor, name string) {
	statuses := []string{"Healthy", "Unhealthy", "Unknown"}
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		status := statuses[rand.Intn(len(statuses))]
		p.UpdateHealth(name, status)
		fmt.Printf("Service %s updated to %s\n", name, status)
	}
}

func main() {
	proctor := NewProctor()

	// Register services
	proctor.RegisterService("ServiceA")
	proctor.RegisterService("ServiceB")
	proctor.RegisterService("ServiceC")

	// Start monitoring
	go proctor.Monitor()

	// Simulate services reporting their health
	go SimulateService(proctor, "ServiceA")
	go SimulateService(proctor, "ServiceB")
	go SimulateService(proctor, "ServiceC")

	// Simulate running for 10 seconds
	time.Sleep(10 * time.Second)

	// Stop monitoring
	proctor.Stop()

	// Print all alerts
	for alert := range proctor.alerts {
		fmt.Println("Logged:", alert)
	}
}
