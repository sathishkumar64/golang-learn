package main

import (
	"fmt"
)


func setmessage(ch chan string){
	ch <- "Hey I'm sending messages to you!"
}

func getmessage(ch chan string){
	msg := <- ch
	fmt.Println(msg)
}

func main(){
	ch := make(chan string)	
	go setmessage(ch)
	go getmessage(ch)

	var input string
	fmt.Scanln(&input)
}
