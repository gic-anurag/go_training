package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main1() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://m001-student:<m001-mongodb-basics>@sandbox.2kbqgtx.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		//panic(err)
		fmt.Println("fatal error")
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			//panic(err)
			fmt.Println("connection close")

		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		//panic(err)
		fmt.Println("if there is an error")

	}
	fmt.Println("Successfully connected and pinged.")
}
