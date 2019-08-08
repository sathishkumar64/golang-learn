package main

import (
	"os"
	"fmt"
)

func main(){
	createFile()
	readFile()
}


func createFile(){
	file,err := os.Create("TestGofile")
	if err != nil{
		panic(err)
	}
	data := []byte("Hello Go world")
	file.Write(data)
	file.WriteString("I'm doing something")
	file.Close()
}

func readFile(){
	file,err := os.Open("TestGofile")
	if err != nil{
		panic(err)
	}

	data := make([]byte, 3)
	
	file.Seek(5,0)
	file.Read(data)
	fmt.Printf("Here is data:: %v" ,string(data))
	file.Close()
}

