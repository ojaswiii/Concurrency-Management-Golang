package main

import (
	"log"
	"net/http"

	"github.com/ojaswiii/MoMoney-Technical-Assignment/internal/driver"
)

func main() {
	// Create and connect to MongoDB client
	driver.ConnectDB()

	// Start the server
	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", routes()))
}
