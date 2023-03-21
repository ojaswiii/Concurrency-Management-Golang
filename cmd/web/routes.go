package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define the route for /posts/:id
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")

	// Start the server
	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
