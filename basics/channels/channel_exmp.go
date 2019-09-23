package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	ch := make(chan int)
	wg.Add(2)

	go func(ch <-chan int) {
		for {

			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}

		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		for i := 0; i < 5; i++ {
			ch <- i

		}
		//ch <- 42
		//	ch <- 64
		//ch <- 100
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
