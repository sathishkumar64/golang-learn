package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	s, err := getDivide(5.0, 6.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func getDivide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
