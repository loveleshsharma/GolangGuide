package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("abc.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)

	// for {
	// 	line, err := r.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		fmt.Println("error reading file")
	// 		break
	// 	}
	// 	fmt.Print(line)
	// }

	for line, err := r.ReadString('\n'); err != io.EOF; line, err = r.ReadString('\n') {

		fmt.Print(line)

	}

}
