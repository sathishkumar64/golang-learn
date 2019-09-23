package main

import (
	"fmt"
	"strconv"
)

func main() {

	var age int
	age = 32

	fmt.Printf("%v, %T", age, age)

	var agenew string

	agenew = strconv.Itoa(age)

	fmt.Printf("%v, %T", agenew, agenew)
}
