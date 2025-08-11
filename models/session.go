package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Session struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string        `bson:"username" json:"username"`
	Token     string        `bson:"token" json:"token"`
	ExpiresAt time.Time     `bson:"expires_at" json:"expires_at"`
}
