package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"fmt"
)

type studentDetails struct {
	StudentID   int    `json:"student_id,omitempty"`
	StudentName string `json:"student_name,omitempty"`
	SchoolName  string `json:"school_name,omitempty"`
	//Address address
}

type address struct{
	Address1 string  `json: Address1`
	City string `json :city`
	State string `json:state`	
}


func main(){

	//createJson()
	readJson()
}

func createJson(){

	students := []studentDetails{
		{
			StudentID:1,
			StudentName:"Sathish",			
			SchoolName: "TNAP",
			/*Address : {
				Address1 : "Kurubarahallai",
				City : "Bangalore",
				State : "KA",
			 },*/
		},
		{
			StudentID:2,
			StudentName:"Babu",			
			SchoolName: "TNAP",
			/*Address : {
				Address1 : "Kurubarahallai",
				City : "Bangalore",
				State : "KA",
			 },*/
			},
	}

	var buf = new (bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(students)
	f ,_ := os.Create("sample.json")	
	io.Copy(f,buf)
	//io.Copy(os.Stdout,buf)
}

func readJson(){ 
	f ,_ := os.Open("sample.json")
	//fmt.Println(f)
	dec := json.NewDecoder(f)

	defer f.Close()
	db := []studentDetails {}
	dec.Decode(&db)
	fmt.Println(db)


}