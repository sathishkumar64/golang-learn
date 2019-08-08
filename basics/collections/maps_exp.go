package main

import "fmt"

func main (){

	//Creating different ways of maps
	/*studentsDetails := map[string] int{
		"Anbu" : 24,
		"Babu" : 32,
		"Chandru" :38,
	}*/

	studentsDetails := make(map[string] int)
	studentsDetails = map[string] int{
		"Anbu" : 24,
		"Babu" : 32,
		"Chandru" :38,
	}
	fmt.Println("Source Student Details::::",studentsDetails)
	studentsDetails["Kamal"]= 39
	fmt.Println("Added new value with other  Student Details:::",studentsDetails)
	fmt.Println("Added new value Student Details:::",studentsDetails["Kamal"])

	delete(studentsDetails ,"Chandru")
	fmt.Println("After deleted Student Details:::",len(studentsDetails))


	// ok values

	//pop,ok := studentsDetails["KamalR"]
	//fmt.Println(pop,ok)

	//_, ok := studentsDetails["Kamal"]
	//fmt.Println(ok)

	//Copy map one to another -- manipulation has happend source as well
	

	sp:= studentsDetails
	delete(sp,"Anbu")
	fmt.Println("After source deleted Student Details:::",studentsDetails)
	fmt.Println("After des deleted Student Details:::",sp)

}
