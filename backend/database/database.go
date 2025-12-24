
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

	// Configure client options with explicit TLS settings
	clientOptions := options.Client().ApplyURI(mongoURL)

	// Set server API version for better compatibility
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions.SetServerAPIOptions(serverAPI)

	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB: ", err)
	}

	log.Println("Successfully connected to MongoDB!")
	return client.Database("mangal_chai_db")
}

func Disconnect() {
	if client != nil {
		client.Disconnect(context.TODO())
	}
}
