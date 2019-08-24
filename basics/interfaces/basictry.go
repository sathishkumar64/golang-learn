package main

import "fmt"

type person struct {
	myname string
}

type getName interface {
	interfaceName() string
}

func (p person) interfaceName() string {
	return p.myname
}

func main() {
	s := person{"Sathish"}
	fmt.Println(s.callFuncName())
}
