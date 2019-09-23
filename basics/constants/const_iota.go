package main

import "fmt"

const (
	a = iota
	b
	c
)

const (
	_ = iota + 9
	a1
	b1
	c1
)

func main() {

	fmt.Println(a, b, c)
	fmt.Println(a1)

}
