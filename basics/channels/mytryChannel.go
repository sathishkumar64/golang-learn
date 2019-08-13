package main

import (
	
	"fmt"
)


func setmessage(ch chan string, frmconsole string){
	for{
		ch <- frmconsole + " Hey I'm sending messages to you!"
		//runtime.Gosched
	}
	
}

func getmessage(ch chan string){
	msg := <- ch
	fmt.Println(msg + " received too")
}

func main(){
	ch := make(chan string)	
	go setmessage(ch,"welcome")
	go getmessage(ch)
	defer fmt.Println("I'm done")
	
}
