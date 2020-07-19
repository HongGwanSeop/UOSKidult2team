package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Person struct {
	name   string `bson:"name"`
	number string `bson:"number"`
}

func create(col *mongo.Collection, name, number string) {
	result, err := col.InsertOne(context.TODO(), bson.D{
		{Key: "name", Value: name},
		{Key: "number", Value: number},
	})
	if err != nil {
		fmt.Println(number)
		return
	}
	fmt.Println(result)
}

func read(col *mongo.Collection, name ...string) {
	var p map[string]string
	var m bson.M
	for _, n := range name {
		filter := bson.D{{Key: "name", Value: n}}
		err := col.FindOne(context.TODO(), filter).Decode(&m)
		bsonBytes, _ := bson.Marshal(m)
		bson.Unmarshal(bsonBytes, &p)
		if err != nil {
			fmt.Println(n)
			fmt.Println(err)
			continue
		}
		fmt.Println("name: ", p["name"], "number: ", p["number"])
	}
}

func update(col *mongo.Collection, name, number string) {
}

func deleteName(col *mongo.Collection, name string) {
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
	case "delete":
	default:
		invalid()
	}
}
