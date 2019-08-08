package main

import (
	"fmt"
	"log"
)

func main (){
/*
	a,b :=1 ,0 

	fmt.Println("Start")
	panic ("Something bad")
	ans := 1/0
	fmt.Println(ans)*/


	/*fmt.Println("Start")
	defer fmt.Println("Middle")  
	panic("Something bad")
	fmt.Println("End")*/

	fmt.Println("Start")
	panicer()
	fmt.Println("End")
}


func panicer(){
	fmt.Println("About to start")

	defer func(){
		if err := recover() ; err !=nil{
			log.Println("Error:",err)
			panic(err)
		}
	}()
	panic("Something went wrong")
	fmt.Println("About to start")
}