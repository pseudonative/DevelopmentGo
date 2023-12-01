package main

import (
	"log"
	"net/http"

	"github.com/pseudonative/my-rest-api/internal/handlers"
)

func main() {
	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)

	// Start server
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
