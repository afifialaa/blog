package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	ID      primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	User    string             `json:"user"`
	Title   string             `json:"title"`
	Content string             `json:"content"`
}