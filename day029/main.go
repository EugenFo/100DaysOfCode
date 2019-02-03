package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	/* client, err := mongo.NewClient("mongodb://localhost:8081")
	if err != nil {
		fmt.Println("error while connecting to database")
	}
	if err == nil {
		fmt.Println("It works!")
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx) */

	ctx, cancle := context.WithTimeout(context.Background(), 2*time.Second)
	fmt.Println("context:", ctx)
	client, err := mongo.Connect(ctx, "mongodb://localhost:8081")
	if err != nil {
		fmt.Println("error", err)
	}
	defer cancle()
	// create new collection instance
	/* newCollection := client.Database("newdb").Collection("testing")
	res, err := newCollection.InsertOne(ctx, bson.M{"name": "some", "value": "thing"})
	fmt.Println("inserted Doc:", res.InsertedID) */

	// query with cursor
	newCollection := client.Database("newdb").Collection("testing")
	cur, err := newCollection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println("error:", err)
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			fmt.Println(err)
		}
		// look up a specific field
		fmt.Println("result:", result["name"])
		fmt.Println("result2:", result)
	}

	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}

	// single Item
	filter := bson.M{"name": "some"}
	var result bson.M
	err = newCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result of 1 item:", result["value"])
}
