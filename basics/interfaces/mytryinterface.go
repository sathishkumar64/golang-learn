package main

import (
	"fmt"
)

type studentDetail struct {
	StudentID   int
	StudentName string
	SchoolName  string
}

type student interface {
	StudentCreate() (studentDetail, error)
}

func (s studentDetail) StudentCreate() (string, error) {
	return s.StudentName, nil
}

func main() {

	student := studentDetail{
		StudentID:   100,
		StudentName: "Sathish",
		SchoolName:  "TNAP",
	}
	fmt.Printf("%v ", student.StudentCreate())
}
