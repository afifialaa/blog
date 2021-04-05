package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/afifialaa/blog/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Comment struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	User       string             `bson:"user" json:"user"`
	Body       string             `bson:"body" json:"body"`
	Posted_At  time.Time          `bson:"posted_at" json:"posted_at"`
	Article_ID primitive.ObjectID `bson:"article_id" json:"article_id,omitempty"`
}

func CreateComment(comment Comment) bool {
	_, err := database.CommentsCollection.InsertOne(context.TODO(), comment)
	if err != nil {
		fmt.Println("mongodb error ", err.Error())
		return false
	}

	fmt.Println("Comment was posted")
	return true
}

func FetchComments(article_id string) []Comment {
	articlePrimID, _ := primitive.ObjectIDFromHex(article_id)
	filter := bson.M{"article_id": articlePrimID}

	// Finding multiple documents returns a cursor
	cursor, err := database.CommentsCollection.Find(context.TODO(), filter)

	if err != nil {
		fmt.Println(err)
	}

	var result []Comment

	// Iterate over the cursor and decode each document
	for cursor.Next(context.TODO()) {
		var comment Comment

		err := cursor.Decode(&comment)

		if err != nil {
			fmt.Println(err)
		}

		result = append(result, comment)
	}

	cursor.Close(context.TODO())
	return result
}

func DeleteComment(id string) bool {
	idPrimitive, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": idPrimitive}

	_, err := database.CommentsCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println("Failed to delete comment")
		fmt.Println(err)
		return false
	}

	return true
}

func UpdateComment(comment Comment) bool {
	filter := bson.M{"_id": comment.ID}
	update := bson.M{
		"$set": bson.M{
			"body": comment.Body,
		},
	}

	_, err := database.CommentsCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
