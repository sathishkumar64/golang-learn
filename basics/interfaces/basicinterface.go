package main

import "fmt"

func main() {

	var r Reader = ConsoleReader{}
	r.Read([]byte("Hello go"))

}

type Reader interface {
	Read([]byte) (int, error)
}

type ConsoleReader struct{}

func (r ConsoleReader) Read(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
