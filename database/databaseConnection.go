package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	Setenv()
	MongoDb := os.Getenv("MongoDb")
	// fmt.Print("DB-URL -> "+MongoDb)

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	// defer cancel()

	fmt.Println("connected to mongodb")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

	return collection
}
