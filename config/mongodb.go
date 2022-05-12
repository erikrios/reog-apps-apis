package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBDatabase() (client *mongo.Client, err error) {
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority", username, password, host, port)

	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	return
}
