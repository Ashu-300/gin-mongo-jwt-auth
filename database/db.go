package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func ConnectDb() {
	url := os.Getenv("MONGO_URL")
	
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal("Error creating Mongo client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

 

	log.Println("✅ Connected to MongoDB")

	db = client.Database("gin_jwt_database")
	
	
}

func GetCollection(name string) *mongo.Collection {
	if db == nil {
		log.Fatal("Database not connected.")
	}
	return db.Collection(name)
}