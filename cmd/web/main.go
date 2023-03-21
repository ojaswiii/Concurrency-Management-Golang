package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const portNumber = ":8080"

func init() {
	// Load the .env file and get database URI
	godotenv.Load("config.env")
	dbURI := os.Getenv("DATABASE")

	// Initialise a MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the MongoDB client
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
}
