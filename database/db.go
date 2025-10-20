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
var userCollection *mongo.Collection

func ConnectDb() {
	url := os.Getenv("MONGO_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	client, err := mongo.Connect(ctx  , options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal("Error creating Mongo client:", err)
	}

	userCollection = client.Database("gin_jwt_database").Collection("user")
	

 

	log.Println("✅ Connected to MongoDB")

	
	
	
}

func GetUserCollection(name string) *mongo.Collection {
	
	return userCollection
}