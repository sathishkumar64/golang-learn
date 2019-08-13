package main

import (
	"fmt"
)


type studentDetail struct {
	StudentID   int   
	StudentName string 
	SchoolName  string 
}

type Student interface{
	StudentCreate(studentDetail struct) (studentDetail, error)
}




func (s studentDetail) StudentCreate() (studentDetail,error){
	student := studentDetail{
				StudentID:1,
				StudentName:"Sathish",			
				SchoolName: "TNAP",	
			}
	return student ,nil
}

func main(){
	var student Students
	student.StudentCreate()
	fmt.Printf("%v  %T" ,student,student)
}