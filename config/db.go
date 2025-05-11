package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var DB *mongo.Client

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(options.Client().ApplyURI("[your-mongodb-address]").SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	fmt.Println("✅ Successfully connected to MongoDB!")

	DB = client
	return DB
}
