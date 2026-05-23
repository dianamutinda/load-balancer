package main

import (
	"fmt"
)
type Server struct{
	Address string
	Health bool
	Name string
}

func (s Server) isHealthy() bool{
	return s.Health
}
func main(){
	s1 := Server{
	Address: "localhost:8080",
	Health: true,
	Name: "server1",
}
	result := s1.isHealthy()
	fmt.Println("Results:", result)
	fmt.Println(s1.Address)
}