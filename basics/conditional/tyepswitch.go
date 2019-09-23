package main

import "fmt"

func main() {

	//var i interface{} =[3]int {}
	var i interface{} = 1

	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		fmt.Println("I will print along with you")
	case float64:
		fmt.Println("i is a float")
		break
		fmt.Println("hey,you dont want print with you")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("not one | two")
	}

}
