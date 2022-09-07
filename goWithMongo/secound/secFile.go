package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"Name,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Podcast     primitive.ObjectID `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Duration    int32              `bson:"duration,omitempty"`
}

func main() {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo DB connection done")
	defer client.Disconnect(context.TODO())
	database := client.Database("quickstart")
	podcastsCollection := database.Collection("podcasts")
	//episodesCollection := database.Collection("episodes")

	podcast := Podcast{
		Name:   "The Developer",
		Author: "anurag singh",
		Tags:   []string{"development", "programming", "coding"},
	}

	result1, err := podcastsCollection.InsertOne(context.TODO(), podcast)
	//result2, err := episodesCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("inserted %v", result1.InsertedID)

	cursor, err := podcastsCollection.Find(context.TODO(), podcast)

	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		// Declare a result BSON object
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			fmt.Println("cursor.Next() error:", err)
			os.Exit(1)
			// If there are no cursor.Decode errors
		} else {
			fmt.Println("\nresult type:", reflect.TypeOf(result))
			fmt.Println("result:", result)
		}

	}
}
