package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "school"
const collecName = "class"

var collection *mongo.Collection

func main() {

	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo DB connection done")

	coll := client.Database("myDB").Collection("favorite_books")
	/*	doc := bson.D{{"title", "Invisible Cities"}, {"author", "Italo Calvino"}, {"year_published", 1974}}
		result, err := coll.InsertOne(context.TODO(), doc)
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)*/

	/*	doc1 := []interface{}{
			bson.D{{"title", "Invisible"}, {"author", "Italo"}, {"year_published", 1987}},
			bson.D{{"title", "kelvin"}, {"author", "uganda"}, {"year_published", 1873}},
			bson.D{{"title", "calin"}, {"author", "rangila"}, {"year_published", 0000}},
		}

		result, err := coll.InsertMany(context.TODO(), doc1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedIDs...)*/

	/*	id, _ := primitive.ObjectIDFromHex("63048c415fe589701c35926d")
		result, err := coll.UpdateOne(
			context.TODO(),
			bson.M{"_id": id},
			bson.D{
				{"$set", bson.D{{"author", "anurag singh"}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)*/

	//Deleting a document

	/*	filter := bson.D{{"author", "anurag singh"}}
		result, err := coll.DeleteOne(context.TODO(), filter)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Deleted %v Documents!\n", result.DeletedCount)
	}*/

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", "Invisible"}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return
		}
		panic(err)
	}
	fmt.Println(result)
}
