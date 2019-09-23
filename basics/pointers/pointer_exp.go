package main

import "fmt"

func main() {

	/*var a int = 64
	var b *int = &a

	fmt.Println(a,b)
	fmt.Println(&a ,b)
	fmt.Println(a,*b)
	*b =65
	fmt.Println(a,*b)*/

	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v,%v,%p \n", a, *b, c)

}
