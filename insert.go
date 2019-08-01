package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func insert(client *mongo.Client, collection *mongo.Collection) {

	persons := createnew()

	insertResult, err := collection.InsertOne(context.TODO(), persons)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
}
func createnew() interface{} {
	var name, dob, home string
	var age int

	fmt.Println("Enter name: ")
	fmt.Scan(&name)
	fmt.Println("Enter age: ")
	fmt.Scan(&age)
	fmt.Println("Enter date of birth (dd-mm-yyyy format): ")
	fmt.Scan(&dob)
	fmt.Println("Enter your hometown: ")
	fmt.Scan(&home)

	people := user{Name: name, Age: age, Dob: dob, City: home}
	return people
}
