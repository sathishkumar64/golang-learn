package main

import "fmt"

func main () {

	//fmt.Println("Start")
	//defer fmt.Println("Middle")  // It will help to get request & response call, but not in no of loops requests.
	//fmt.Println("End")


	a := "Sat"
	defer fmt.Println(a)
	a ="kumar"
}