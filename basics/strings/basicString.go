package main

import "fmt"

func main() {
	s := "this is string"

	fmt.Printf("%v,%T", s, s)

	//Want to get 2 index

	fmt.Printf("%v,%T", s[2], s[2])

	//want to get bytes
	b := []byte(s)

	fmt.Printf("%v,%T", b, b)
}
