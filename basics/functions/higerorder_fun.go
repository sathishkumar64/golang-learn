package main

import "fmt"

func main(){
	sum(5,num(5))
}

func num(n int) int{
	return n *2
}

func sum(a int, f int) {
	total := a + f
	fmt.Println(total)
}