package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ojaswiii/MoMoney-Technical-Assignment/internal/handlers"
)

func routes() http.Handler {
	// Create a new router
	router := mux.NewRouter()

	// Define the route for /posts/:id
	router.HandleFunc("/posts/{id}", handlers.GetPost).Methods("GET")

	return router
}
