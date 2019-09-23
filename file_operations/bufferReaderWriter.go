package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	createFile()
	readFile()
}

func createFile() {
	file, err := os.Create("TestGofile_buf")
	if err != nil {
		panic(err)
	}
	buf := bufio.NewWriter(file)
	buf.WriteString("This is Buffer Writer Apporach")
	buf.Flush()
	file.Close()
}

func readFile() {
	file, err := os.Open("TestGofile_buf")
	if err != nil {
		panic(err)
	}

	buf := bufio.NewReader(file)
	data, _ := buf.Peek(5)
	fmt.Printf("Here is data:: %v", string(data))
	file.Close()
}
