package main

import (
	"context"
	"fmt"
	"log"

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
	//replace <username with user id>
	//replace <password with database password>
	//after replacing <remove angular brackets>

	clientOptions := options.Client().ApplyURI("mongodb+srv://<username>:<password>@db-xhiqs.mongodb.net/test")
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
