package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("error", err)
	}
	defer file.Close()

	buffer := make([]byte, 1)

	_, err = file.Read(buffer)
	fmt.Println(string(buffer))

	// _, err = file.Read(buffer)
	// fmt.Println(string(buffer))

	// _, err = file.Read(buffer)
	// fmt.Println(string(buffer))

	// _, err = file.Read(buffer)
	// fmt.Println(string(buffer))

	// var prefix = []byte("abcd\r\npqr\r\n")
	// var a = []byte("pqr\r\n")
	// fmt.Println("b:", a)

	// lines := bytes.Split(prefix, []byte{'\n'})
	// fmt.Println("len", len(lines))
	// fmt.Println(string(lines[len(lines)-1]))
	// fmt.Println("aft:", lines[len(lines)-2])
	// if bytes.Contains(lines[len(lines)-2], []byte{'\n'}) {
	// 	fmt.Println("found!")
	// }

	// if bytes.Contains(prefix, make([]byte, len(prefix))) || bytes.Contains(prefix, []byte{'\n'}) {
	// 	fmt.Println("inside")
	// }

	// buff := make([]byte, 10)
	// fmt.Println("length: ", cap(buff))

	// buff = buff[:0]
	// fmt.Println("length: ", cap(buff))

	var b byte

	fmt.Print(b)

	arr := []byte("jdshjkds")
	fmt.Println(arr)

	for i := range arr {
		arr[i] = b
	}
	fmt.Println(arr)

	file.Read
}
