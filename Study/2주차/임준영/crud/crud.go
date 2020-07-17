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
	var p Person
	var d bson.D
	for _, n := range name {
		err := col.FindOne(context.TODO(), bson.D{
			{Key: "name", Value: n},
		}).Decode(&d)
		if err != nil {
			fmt.Println(n)
			fmt.Println(err)
			continue
		}
		/// p.name = d.Map().(primitive.M)["name"]
		fmt.Println("name: ", p.name, "number: ", p.number)
		fmt.Printf("%+v\n", p)
		fmt.Printf("%+v\n", d)
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
