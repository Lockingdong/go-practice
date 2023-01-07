package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostDoc struct {
	Title      string    `bson:"title"`
	Content    string    `bson:"content"`
	DisabledAt time.Time `bson:"disabled_at"`
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	postCollection := client.Database("test").Collection("post")

	// Find
	docs := make([]PostDoc, 0)

	findOption := options.Find()
	findOption.SetProjection(bson.D{
		{Key: "title", Value: 1},
		{Key: "content", Value: 1},
		{Key: "disabled_at", Value: 1},
	})
	findOption.SetSort(bson.D{{Key: "disabled_at", Value: 1}})
	findOption.SetLimit(1)
	findOption.SetSkip(1)
	curr, err := postCollection.Find(context.TODO(), bson.M{"title": "test_titl"}, findOption)
	if err != nil {
		fmt.Println(err)
	}
	if err := curr.All(context.TODO(), &docs); err != nil {
		fmt.Println(err)
	}
	fmt.Println(docs) // [{ test_title updated_test_content 0001-01-01 00:00:00 +0000 UTC} ...

	// CountDocuments
	count, err := postCollection.CountDocuments(context.Background(), bson.M{"title": "test_title"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count) // 32

	// FindOne
	var doc PostDoc
	findOneOption := &options.FindOneOptions{}
	findOneOption.SetProjection(bson.D{{Key: "title", Value: 1}})
	result := postCollection.FindOne(context.TODO(), bson.M{"title": "test_title"}, findOneOption)
	if err := result.Decode(&doc); err != nil {
		fmt.Println(err) // mongo: no documents in result
	} else {
		fmt.Println(doc) // {test_title  0001-01-01 00:00:00 +0000 UTC}
	}

	// FindOneAndUpdate
	findOneAndUpdateOptions := options.FindOneAndUpdate()
	findOneAndUpdateOptions.SetUpsert(false)
	findOneAndUpdateOptions.SetReturnDocument(options.After)

	update := bson.M{
		"$set": bson.M{"content": "find one and update demo"},
	}

	result = postCollection.FindOneAndUpdate(
		context.TODO(),
		bson.M{"title": "test_title"},
		update,
		findOneAndUpdateOptions,
	)
	if err := result.Decode(&doc); err != nil {
		fmt.Println(err) // return "mongo: no documents in result" if SetUpsert(false)
	} else {
		fmt.Println(doc) // {test_title find one and update demo 0001-01-01 00:00:00 +0000 UTC}
	}
}
