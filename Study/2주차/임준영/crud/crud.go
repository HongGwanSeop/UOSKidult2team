package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID     primitive.ObjectID `bson:"_id"`
	name   string             `bson:"name"`
	number string             `bson:"number"`
}

func create(col *mongo.Collection, name, number string) {
	_, err := col.InsertOne(context.TODO(), bson.D{
		{Key: "name", Value: name},
		{Key: "number", Value: number},
	})
	if err != nil {
		fmt.Println(number)
		return
	}
	fmt.Printf("Create Success: <%s, %s>\n", name, number)
}

func read(col *mongo.Collection, name ...string) {
	var p map[string]string
	for _, n := range name {
		filter := bson.D{{Key: "name", Value: n}}
		err := col.FindOne(context.TODO(), filter).Decode(&p)
		if err != nil {
			fmt.Printf("Read Fail: <%s>, ", n)
			fmt.Println(err)
			continue
		}
		fmt.Println("name: ", p["name"], "number: ", p["number"])
		// fmt.Printf("%+v\n", p)
	}
}

func update(col *mongo.Collection, name, number string) {
	filter := bson.D{{Key: "name", Value: name}}
	upd := bson.D{{"$set", bson.D{{Key: "number", Value: number}}}}
	result, err := col.UpdateOne(context.TODO(), filter, upd, options.Update().SetUpsert(true))
	if err != nil {
		fmt.Println(err)
		return
	}
	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
		return
	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
		return
	}
}

func deleteName(col *mongo.Collection, name string) {
	result, err := col.DeleteMany(context.TODO(), bson.D{
		{Key: "name", Value: name},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if result.DeletedCount > 0 {
		fmt.Printf("Delete Success: <%s>\n", name)
	} else {
		fmt.Printf("Delete Fail: <%s>\n", name)
	}
}

func crud(col *mongo.Collection, args ...string) {
	switch args[0] {
	case "create":
		if len(args) != 3 {
			invalid()
			return
		}
		create(col, args[1], args[2])
	case "read":
		if len(args) == 1 {
			invalid()
			return
		}
		read(col, args[1:]...)
	case "update":
		if len(args) != 3 {
			invalid()
			return
		}
		update(col, args[1], args[2])
	case "delete":
		if len(args) != 2 {
			invalid()
			return
		}
		deleteName(col, args[1])
	default:
		invalid()
	}
}
