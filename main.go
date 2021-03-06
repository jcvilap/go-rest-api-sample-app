package main

import (
	"net/http"
	"route"
	"log"
)

func main() {
	// Create routes
	route.Route()
	
	// Start server
	err := http.ListenAndServe(":8080", nil)
	
	// Log any error
	log.Fatal(err)
}