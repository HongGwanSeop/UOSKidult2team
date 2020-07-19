package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func invalid() {
	fmt.Println("Invalid argument\nex) -------------------------------------\n$ ./crud create <name> <number>\n$ ./crud read <name1> <name2> <name3> ...\n$ ./crud update <name> <newnumber>\n$ ./crud delete <name>")
}

func mongoConn() (client *mongo.Client) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return client
}

func main() {
	if len(os.Args) == 1 {
		invalid()
		return
	}
	args := os.Args[1:]
	conn := mongoConn()
	if conn == nil {
		return
	}
	col := conn.Database("test").Collection("people")
	crud(col, args...)
}
