package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func main() {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("Error creating Consul client:", err)
		return
	}

	registration := &api.AgentServiceRegistration{
		ID:      "user-service-1",
		Name:    "user-service",
		Port:    8080,
		Address: "localhost",
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		fmt.Println("Error registering service:", err)
	}
}
