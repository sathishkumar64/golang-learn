package main

import "fmt"

func main() {

	/*switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default :
		fmt.Println("not one | two")
	}
	*/

	// Multiple test cases
	/*switch 10 {
	case 1,5,10:
		fmt.Println("one, five, ten")
	case 2,4,6:
		fmt.Println("two,four,six")
	default :
		fmt.Println("not one | two")
	}*/

	// Multiple test cases with condition
	/*switch i:= 2+3;i {
	case 1,5,10:
		fmt.Println("one, five, ten")
	case 2,4,6:
		fmt.Println("two,four,six")
	default :
		fmt.Println("not one | two")
	}*/

	// tag less test
	i := 10
	switch {
	case i <= 10:
		fmt.Println("one, five, ten")
		fallthrough /// user control for this switch statement
	case i <= 20:
		fmt.Println("two,four,six")
	default:
		fmt.Println("not one | two")
	}
}
