package driver

import (
	"context"
	"log"

	"github.com/ojaswiii/MoMoney-Technical-Assignment/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() *mongo.Client {
	// // Load the .env file and get database URI
	// godotenv.Load("config.env")
	// dbURI := os.Getenv("DATABASE")

	// Initialize MongoDB client
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://ojaswisaxena14:FlSU8R3Cmy8hbBV9@cluster0.l30ec2b.mongodb.net/formo?retryWrites=true&w=majority"))
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

func FindPost(id int) error {
	// Check if the post is already in the database
	collection := client.Database("formo").Collection("posts")
	filter := bson.M{"id": id}
	var post models.Post
	return collection.FindOne(context.Background(), filter).Decode(&post)
}

func SavePost(post models.Post) {
	collection := client.Database("formo").Collection("posts")
	// Save the post in the database
	_, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		log.Println("Failed to save post to database:", err)
	}
}
