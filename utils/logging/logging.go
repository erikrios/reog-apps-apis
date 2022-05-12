package logging

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Logging interface {
	Error(message string)
}

type mongoLogging struct {
	client *mongo.Client
}

func NewMongoLogging(client *mongo.Client) *mongoLogging {
	return &mongoLogging{client: client}
}

func (m *mongoLogging) Error(message string) {
	coll := m.client.Database("logging").Collection("errors")
	doc := bson.D{{"message", message}, {"time", time.Now().UnixNano()}}

	if id, err := coll.InsertOne(context.Background(), doc); err != nil {
		log.Println(err)
	} else {
		log.Println(id)
	}
}
