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
	ID         string    `bson:"id"`
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

	// InsertOne
	doc := &PostDoc{
		Title:   "test_title",
		Content: "test_content",
	}
	insertOneRes, _ := postCollection.InsertOne(context.TODO(), doc)
	fmt.Println(insertOneRes.InsertedID)
	// ObjectID("63b9347af88f4e2d1367fe40")

	// InsertMany
	docs := []any{
		&PostDoc{
			Title:   "test_title",
			Content: "test_content",
		},
		&PostDoc{
			Title:   "test_title",
			Content: "test_content",
		},
	}
	insertManyRes, _ := postCollection.InsertMany(context.TODO(), docs)
	fmt.Println(insertManyRes.InsertedIDs...)
	// ObjectID("63b936e447d074a9b9cf500f") ObjectID("63b936e447d074a9b9cf5010")

	// UpdateOne and UpdateMany
	filter := bson.M{
		"title": "test_title",
	}
	update := bson.M{
		"$set": bson.M{"content": "updated_test_content"},
	}
	updateRes, _ := postCollection.UpdateOne(context.TODO(), filter, update)
	fmt.Println(updateRes.MatchedCount)  // 1
	fmt.Println(updateRes.ModifiedCount) // 1
	fmt.Println(updateRes.UpsertedCount) // 0
	fmt.Println(updateRes.UpsertedID)    // <nil>

	updateRes, _ = postCollection.UpdateMany(context.TODO(), filter, update)
	fmt.Println(updateRes.MatchedCount)  // 17
	fmt.Println(updateRes.ModifiedCount) // 16
	fmt.Println(updateRes.UpsertedCount) // 0
	fmt.Println(updateRes.UpsertedID)    // <nil>

	// Update with Upsert
	opdateOpt := &options.UpdateOptions{}
	opdateOpt.SetUpsert(true)
	upsertFilter := bson.M{
		"title": "upsert_title",
	}
	upsertUpdate := bson.M{
		"$set": bson.M{
			"content": "upsert_content",
		},
	}
	updateRes, _ = postCollection.UpdateOne(context.TODO(), upsertFilter, upsertUpdate, opdateOpt)
	fmt.Println(updateRes.MatchedCount)  // 0
	fmt.Println(updateRes.ModifiedCount) // 0
	fmt.Println(updateRes.UpsertedCount) // 1
	fmt.Println(updateRes.UpsertedID)    // ObjectID("63b93da96cff122ceddefca0")
}
