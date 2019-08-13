package main

import (
	"fmt"
)

type Students interface{
	StudentCreate() (string,error)
}

type studentDetail struct {}


func (s studentDetail) StudentCreate() (string,error){
	return "Poda dey neeyum unkodingum" ,nil
}

func main(){
	var student Students
	student.StudentCreate()
	fmt.Printf("%v  %T" ,student,student)
}