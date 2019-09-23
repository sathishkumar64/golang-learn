package main

import "fmt"

func main() {

	if true {
		fmt.Println("The test is true")
	}

	studentsDetails := map[string]int{
		"Anbu":    24,
		"Babu":    32,
		"Chandru": 38,
	}

	if pop, ok := studentsDetails["Babu"]; ok {
		fmt.Println(pop)
	}

	fmt.Println(10 <= 20, 10 >= 30, 10 != 20)

	fmt.Println(1 > 5 || returnTrue())
}

func returnTrue() bool {
	fmt.Println(true)
	return true
}
