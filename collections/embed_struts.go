package main

import "fmt"

type Animal struct{
	Name string
	Origin string
}

type Bird struct{
	Animal
	SpeedKPH float32
	Canfly bool
}

func main (){

/*	b:= Bird{}
	b.Name ="Emu"
	b.Origin ="Australia"
	b.SpeedKPH =48
	b.Canfly=false
*/

b:= Bird{ Animal : Animal{ Name :"Emu", Origin :"Australia"} ,SpeedKPH :48,Canfly:false}

	fmt.Println(b)

	fmt.Println(b.Name)

}
