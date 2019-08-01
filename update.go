package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func update(client *mongo.Client, collection *mongo.Collection) {
	var objID string
	var people user
	fmt.Println("Enter the id of the person: ")
	fmt.Scan(&objID)
	id, _ := primitive.ObjectIDFromHex(objID)
	result := collection.FindOne(context.TODO(), bson.M{"_id": id})
	result.Decode(&people)
	if people.Age != 0 {
		fmt.Println(people, " found successfully")
		people = updaterequire(client, collection, people)
		collection.FindOneAndUpdate(context.TODO(), bson.M{"_id": id}, bson.M{"$set": people})
		fmt.Println(people, "updated successfully")
	} else {
		fmt.Println("Enter valid id")
	}
}

func updaterequire(client *mongo.Client, collection *mongo.Collection, person user) user {
	var name, dob, home string
	var age int
	var field string

	fmt.Println("Enter the field you want to update: \n Name \n Age \n dob \n home")
	fmt.Println("Enter one of the above fields: ")
	fmt.Scan(&field)

	switch field {
	case "name":
		fmt.Print("Enter new name: ")
		fmt.Scan(&name)
		person.Name = name
	case "age":
		fmt.Print("Enter new age: ")
		fmt.Scan(&age)
		person.Age = age
	case "dob":
		fmt.Print("Enter new dob (in dd-mm-yyyy format): ")
		fmt.Scan(&dob)
		person.Dob = dob
	case "home":
		fmt.Print("Enter new home: ")
		fmt.Scan(&home)
		person.City = home

	default:
		fmt.Println("enter valid values")
	}
	return person
}
