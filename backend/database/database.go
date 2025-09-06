
package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

func Connect() *mongo.Database {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		mongoURL = "mongodb://localhost:27017"
	}

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("mangal_chai_db")
}

func Disconnect() {
	if client != nil {
		client.Disconnect(context.TODO())
	}
}
