package database

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
	"fmt"
	"log"
)

var ArticlesCollection *mongo.Collection
var CommentsCollection *mongo.Collection

func Connect() {
	clientOptions := options.Client().ApplyURI(os.Getenv("CLOUD_MONGO"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	ArticlesCollection = client.Database("blog").Collection("articles")
	CommentsCollection = client.Database("blog").Collection("comments")

}
