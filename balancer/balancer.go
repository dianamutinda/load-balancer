package main

import (
	"fmt"
	"sync"
	"errors"
)
type Server struct{
	Address string
	Health bool
	Name string
}

type LoadBalancer struct{
	servers []Server
	current int
	mu sync.Mutex
}

func NewLoadBalancer (servers []Server) *LoadBalancer{
	return &LoadBalancer{
		servers: servers,
		current: 0,
	}
}
func (lb *LoadBalancer) NextServer() (*Server, error){
	lb.mu.Lock()
    defer lb.mu.Unlock()

	for i := 0; i < len(lb.servers); i++{
	    server := &lb.servers [lb.current]
	    lb.current = (lb.current + 1) % len(lb.servers)
	    
		if server.Health{
			return server, nil
		}
		fmt.Printf("%s is not healthy, skipping\n", server.Name)
	}
	return nil, errors.New("No Healthy servers available")
}
func main(){
	servers := []Server{
		{Address: "localhost: 8080", Health: true, Name: "Server1" },
		{Address: "localhost: 8081", Health: false, Name: "Server2" },
		{Address: "localhost: 8082", Health: true, Name: "Server3" },
	}
	lb := NewLoadBalancer(servers)

	for i := 1; i <= 7; i++{
		server, err := lb.NextServer()
		if err != nil {
			fmt.Println("Error", err)
			continue
		}
		fmt.Printf("Task %d -> assigned to %s \n", i, server.Name)
	}
}