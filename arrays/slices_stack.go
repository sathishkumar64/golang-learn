package main

import "fmt"

func main (){

	a:= [] int {1,2,3,4,5}

	//To remove 3
	b:= a[:2]
	fmt.Printf("Remaining Value %v\n",b)

	//To remove last value
	c:= a[:len(a)-1]
	fmt.Printf("Remaining Value %v\n",c)

	//To remove mniddle value
	d:= append(a[:2],a[3:]...)
	fmt.Printf("Remaining Value %v\n",d)
	fmt.Printf("source Value %v\n",a)
}

