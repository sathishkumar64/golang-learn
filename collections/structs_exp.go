package main

import "fmt"

type Student struct{
	studentId int
	studentName string
	degree []string
}

func main (){
	aStudent := Student{
		studentId :1,
		studentName : "Sathish",
		degree:[]string {
			"ABC","DEF","GH"},
	}

	fmt.Printf("The Student is: %v\n",aStudent)
	fmt.Printf("The Student is: %v, %v\n",aStudent.studentName,aStudent.degree[1])

	aEmp:=struct{name string}{name:"Sat"}
	fmt.Printf("Annomus Sturt %v",aEmp)

	bEmp:=&aEmp
	bEmp.name ="Kumar"
	fmt.Printf("Annomus Sturt %v",aEmp)
	fmt.Printf("Annomus Sturt %v",bEmp)
}


