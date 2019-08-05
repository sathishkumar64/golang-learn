package main

import "fmt"

func main(){

	g := student{
		stuName : "Sathish",
		age : 34,
	}

	g.callStudent()

}


type  student struct {
	stuName string
	age int
}

func (s student ) callStudent(){
	fmt.Println(s.stuName,s.age)
}