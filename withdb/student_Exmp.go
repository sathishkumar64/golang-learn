package main

import (
	"fmt"
)

/* This is Student Object */
type Student struct {
	ID          primitive.ObjectID
	studentID   string
	studentName string
	className   string
	schoolname  string
}

var client *mongo.Client

func main() {
	fmt.Println("Applicaiton is starting...")
}
