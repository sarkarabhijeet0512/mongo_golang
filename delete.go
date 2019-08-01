package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func delete(client *mongo.Client, collection *mongo.Collection) {
	var objID string
	var person user
	fmt.Println("enter id")
	fmt.Scan(&objID)
	id, _ := primitive.ObjectIDFromHex(objID)
	result := collection.FindOneAndDelete(context.TODO(), bson.M{"_id": id})
	result.Decode(&person)
	if person.Age != 0 {
		fmt.Println(person, " deleted successfully")
	} else {
		fmt.Println("delete failed wrong id")

	}
}
