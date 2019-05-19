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

	buffer := make([]byte, 5)
	file.Read(buffer)
	fmt.Println(string(buffer))

	file.Seek(-1, 1)

	var a int64 = 0

	// file.Read(buffer)
	// fmt.Println(string(buffer))

	// file.ReadAt(buffer, 3)
	// fmt.Println(string(buffer))

	// file.ReadAt(buffer, 7)
	// fmt.Println(string(buffer))

	// file.ReadAt(buffer, 3)
	// fmt.Println(string(buffer))

	// str := "thoughtworks university"
	// var prefix []byte

	// prefix = []byte(str)
	// fmt.Println(len(prefix))

	// str = "thequickbrownfoxjumpsoverthelazydog"
	// prefix = []byte(str)
	// fmt.Println(len(prefix))

}
