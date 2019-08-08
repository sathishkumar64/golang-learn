package main

import (
	"runtime"
	"fmt"
	"sync"
)




var wg= sync.WaitGroup{}

/*func  main(){
	msg := "Calling Go Routine"
	wg.Add(1)
	go func(msg string){
		fmt.Println(msg)
		wg.Done()
	}(msg)
	wg.Wait()	
}*/

var counter =0
func main(){

	
	fmt.Printf("Threads:  %v\n",runtime.GOMAXPROCS(-1))
	for i:=0 ;i <10;i++{
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()

}


func sayHello(){
 fmt.Printf("Hello %v\n",counter)
 wg.Done()
}

func increment(){
 counter++
 wg.Done()
}