package main

import "fmt"

func main() {
	name := "sathish"
	age := 34
	height := 5.11
	isSingle := false

	fmt.Println(name, age, height, isSingle)

	// Multi Short declaration

	city, state := "bangalore", "KA"

	fmt.Println(city, state)

	// Reassignment

	name = "kumar"

	fmt.Println(name)
	//same as the above
	name, brith := "sathish vasu", 1984

	fmt.Println(name, brith)
}
