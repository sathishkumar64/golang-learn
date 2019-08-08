package main

import "fmt"

func main() {
	fmt.Println("I've started to learn GOlang...")
	learnPackage()
}


/** Cannot declare same name as again. */
func learnPackage() {
	fmt.Println("Learning Package stuffs")
}
