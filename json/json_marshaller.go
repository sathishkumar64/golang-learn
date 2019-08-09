package main

import (
	"time"
	"encoding/json"
	"fmt"
	"strings"
	"log"
)

type ShirtSize byte

type Person struct {
	Name string    `json:"name,omitempty"`
	Born time.Time `json:"born,omitempty"`
	Size ShirtSize `json:"size,omitempty"`
}

const (
	NA ShirtSize = iota
	XS
	S
	M
	L
	XL
)

func main(){
	
	var p Person
	dec := json.NewDecoder(strings.NewReader(input))

	if err := dec.Decode(&p); err !=nil {
		log.Fatalf("parse person %v",err)
	}
	fmt.Println(p)
}

func ( p *Person) Parse(s string) error{

	return err
}