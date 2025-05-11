package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Book struct {
	ID     bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Title  string        `bson:"title,omitempty" json:"title"`
	Author string        `bson:"author,omitempty" json:"author"`
}
