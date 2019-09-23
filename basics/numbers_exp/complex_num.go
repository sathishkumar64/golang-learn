package main

import "fmt"

func main() {

	var a complex64 = 1 + 2i

	fmt.Printf("%v,%T", a, a)

	fmt.Printf("%v,%T", real(a), real(a))

	fmt.Printf("%v,%T", imag(a), imag(a))
}
