package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("error")
	}
	defer file.Close()

	buffer := make([]byte, 1024)

	n, err := file.Read(buffer)

	if err == io.EOF {
		return
	} else if err != nil {
		fmt.Println("error reading file")
		return
	}
	fmt.Print(string(buffer))
	fmt.Printf("\nno of bytes read: %d", n)

}
