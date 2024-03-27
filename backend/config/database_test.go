package config

import (
	"os"
	"testing"
)

func TestGetMongoClient(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")

	client, err := GetMongoClient()
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	if client == nil {
		t.Errorf("Unexpected client: got nil")
	}
}

func TestGetDatabase(t *testing.T) {
	os.Setenv("MONGO_DATABASE_NAME", "testdb")

	db, err := GetDatabase()
	if err != nil {
		t.Fatalf("Failed to get database: %v", err)
	}

	if db.Name() != "testdb" {
		t.Errorf("Unexpected database name: got %v, expected %v", db.Name(), "testdb")
	}
}
