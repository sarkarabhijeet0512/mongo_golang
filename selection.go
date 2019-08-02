package main

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func db(client *mongo.Client) {
	//replace <database name>
	//replace <collection name>
	//remove <angular brackets>

	collection := client.Database("<database name>").Collection("<collection name>")
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
