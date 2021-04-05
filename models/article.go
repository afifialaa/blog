package models

import (
	"context"
	"fmt"
	"log"

	"github.com/afifialaa/blog/database"
	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Article struct {
	ID      primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	User    string             `bson:"user" json:"user"`
	Title   string             `bson:"title" json:"title"`
	Content string             `bson:"content" json:"content"`
}

func UpdateArticle(article Article) bool {
	data := structs.Map(article)

	_, err := database.ArticlesCollection.UpdateOne(context.TODO(), bson.M{"_id": article.ID},
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

func CreateArticle(article Article) bool {
	_, err := database.ArticlesCollection.InsertOne(context.TODO(), article)
	if err != nil {
		fmt.Println("mongodb error ", err.Error())
		return false
	}

	fmt.Println("Article was created")
	return true
}

func DeleteArticle(id string) bool {
	idPrimitive, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": idPrimitive}

	_, err := database.ArticlesCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println("Failed to delete article")
		return false
	}

	return true
}

func FetchArticles(user string) []Article {
	filter := bson.M{"user": user}
	// Finding multiple documents returns a cursor
	cursor, err := database.ArticlesCollection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		var myslice []Article
		return myslice
	}

	var result []Article

	// Iterate over the cursor and decode each document
	for cursor.Next(context.TODO()) {
		var article Article

		err := cursor.Decode(&article)

		if err != nil {
			fmt.Println(err)
		}

		result = append(result, article)
	}

	cursor.Close(context.TODO())
	return result
}
