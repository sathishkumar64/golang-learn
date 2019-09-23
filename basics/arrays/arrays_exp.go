package main

import "fmt"

func main() {

	grades := [3]int{
		97, 85, 93,
	}

	stuGrades := [...]int{97, 85, 64}
	fmt.Printf("Grades: %v, %v \n", grades, stuGrades)

	var students [3]string
	students[0] = "Arun"
	students[1] = "babu"
	students[2] = "Chan"
	fmt.Printf("Students %v, %v \n", students[2], len(students))

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 1}
	identityMatrix[1] = [3]int{1, 0, 0}
	identityMatrix[2] = [3]int{0, 0, 1}

	//identityMatrix[2]=[4]int{0,0,1,1}  // This is error,based on your declartion

	fmt.Printf("identityMatrix::: %v", identityMatrix)

	a := [...]int{1, 2, 3}
	b := a
	b[1] = 4
	fmt.Println(a, b)

}
