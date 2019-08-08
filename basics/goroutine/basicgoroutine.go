package main

import "fmt"
import  "time"

func  main (){
	go callGo()
	time.Sleep(100 * time.Millisecond)
}

func callGo(){
	fmt.Println("Calling Go Routine")
}