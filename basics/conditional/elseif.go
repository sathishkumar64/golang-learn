package main

import (
	"fmt"
	"math"
)

func main() {
	num := 0.12

	if num == math.Pow(math.Sqrt(num), 2) {
		fmt.Println("These are same")
	} else {
		fmt.Println("These are different")
	}
}
