package config

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbInstance () *mongo.Client{
	MongoDb := "mongodb://localhost:5000"
	fmt.Println(MongoDb)

	client,error := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if error != nil {
		log.Fatal(error)
	}
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)

	defer cancel()

	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB established")
	return client
}

var Client *mongo.Client = DbInstance()

func OpenCollection (client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("clockbuster").Collection(collectionName) 
	return collection
}