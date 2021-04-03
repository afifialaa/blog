package database

import (
	"os"

	"github.com/afifialaa/blog/models"
	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
	"fmt"
	"log"
)

var ArticlesCollection *mongo.Collection

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

}

func UpdateArticle(article models.Article) bool {
	data := structs.Map(article)

	_, err := ArticlesCollection.UpdateOne(context.TODO(), bson.M{"_id": article.ID},
		bson.D{
			{"$set", data},
		},
	)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func CreateArticle(article models.Article) bool {
	_, err := ArticlesCollection.InsertOne(context.TODO(), article)
	if err != nil {
		fmt.Println("mongodb error ", err.Error())
		return false
	}

	fmt.Println("Article was created")
	return true
}

func DeleteArticle(id string) bool {
	fmt.Println(id)
	idPrimitive, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{"_id", idPrimitive}}

	_, err := ArticlesCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println("Failed to delete article")
		return false
	}

	return true
}

func FetchArticles(user string) []models.Article {
	filter := bson.D{{"user", user}}
	// Finding multiple documents returns a cursor
	cursor, err := ArticlesCollection.Find(context.TODO(), filter)

	if err != nil {
		fmt.Println(err)
	}

	var result []models.Article

	// Iterate over the cursor and decode each document
	for cursor.Next(context.TODO()) {
		var book models.Article

		err := cursor.Decode(&book)

		if err != nil {
			fmt.Println(err)
		}

		result = append(result, book)
	}

	cursor.Close(context.TODO())
	return result
}
