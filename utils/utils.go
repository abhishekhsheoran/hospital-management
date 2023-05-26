package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Connection *mongo.Client
)

func InitializeDatabase() {
	options := options.Client().ApplyURI("mongodb://localhost:27017/")
	var err error
	Connection, err = mongo.Connect(context.Background(), options)
	if err != nil {
		log.Fatal(err)
	}
}
