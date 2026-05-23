package main

import (
	"fmt"
	"net/http"
	"log"
)
func handleRequest(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Load balancer is running")
}
func main(){
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}