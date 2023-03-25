package driver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ojaswiii/MoMoney-Technical-Assignment/internal/models"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() *mongo.Client {
	// Load the .env file and get database URI
	viper.SetConfigName("config")
	viper.AddConfigPath("../../")
	er := viper.ReadInConfig()
	if er != nil {
		panic(fmt.Errorf("fatal error config file: %w", er))
	}
	dbURI := (viper.Get("DATABASE")).(string)

	// Initialize MongoDB client
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal(err)
	}

	// Connect to MongoDB
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to db")

	return client
}

func FindPost(id int, w http.ResponseWriter) bool {
	var flag bool = false

	// Check if the post is already in the database
	collection := client.Database("formo").Collection("posts")
	filter := bson.M{"id": id}
	var post models.Post
	err := collection.FindOne(context.Background(), filter).Decode(&post)
	if err == nil {
		// If the post is already in the database, return it
		json.NewEncoder(w).Encode(post)
		flag = true
		log.Println("Post found in database, no need for api call")
	}
	return flag
}

func SavePost(post models.Post) {
	collection := client.Database("formo").Collection("posts")
	// Save the post in the database
	_, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		log.Println("Failed to save post to database:", err)
	}
}
