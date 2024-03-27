package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetMongoClient connects to the MongoDB server and returns a client instance.
func GetMongoClient() (*mongo.Client, error) {
	connectionString := os.Getenv("MONGO_CONNECTION_STRING")
	if connectionString == "" {
		connectionString = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}

func GetDatabase() (*mongo.Database, error) {
	databaseName := os.Getenv("MONGO_DATABASE_NAME")
	if databaseName == "" {
		databaseName = "loginapp"
	}

	client, err := GetMongoClient()
	if err != nil {
		return nil, err
	}
	return client.Database(databaseName), nil
}
