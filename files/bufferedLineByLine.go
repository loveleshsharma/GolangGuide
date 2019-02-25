package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("files/test.txt")
	if err != nil {
		fmt.Println("error", err)
	}
	defer file.Close()

	buffer := make([]byte, 1024)

	_, err = file.Read(buffer)

	if err == io.EOF {
		return
	} else if err != nil {
		fmt.Println("error reading file")
		return
	}
	var lines = bytes.Split(buffer, []byte("\n"))

	for i := 0; i < len(lines); i++ {
		fmt.Println(string(lines[i]))
	}

	//lastNewLineIndex := bytes.LastIndex(buffer, []byte("\r\n"))

	//buffer = append(buffer[lastNewLineIndex:lastNewLineIndex+1], buffer[lastNewLineIndex+1:]...)

	//fmt.Println("\nbytes: ", string(buffer))

	// fmt.Printf("\nno of bytes read: %d", n)

}
