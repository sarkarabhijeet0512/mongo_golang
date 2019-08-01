package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type user struct {
	Name string
	Age  int
	City string
	Dob  string
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://Abhi_0512:12345ok@db-xhiqs.mongodb.net/test")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	db(client)

}

func db(client *mongo.Client) {
	collection := client.Database("mydb").Collection("persons")
	var choice string
	fmt.Println("what operation do you want to perform (insert/delete/update/get/exit)")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid format;", err)
		os.Exit(0)
	}

	switch choice {
	case "insert":
		insert(client, collection)

	case "delete":
		delete(client, collection)

	case "update":
		update(client, collection)

	case "get":
		search(client, collection)
	case "exit":
		os.Exit(0)

	default:
		fmt.Println("invalid input\nuse appropriate values")
	}
}
