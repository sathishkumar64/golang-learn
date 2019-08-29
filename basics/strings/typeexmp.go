package main

import "fmt"

type myage int

func (a myage) getResult() string {
	return fmt.Sprintf("%v",a)
}

type urage struct{
	myage
}

func main(){
a := myage(16*2)
fmt.Println(a)

y := urage{a}
fmt.Println(y)

}

	

