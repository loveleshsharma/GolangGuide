package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	// buffer := []byte("new data in the file!")

	file, err := os.Create("myfile")
	if err != nil {
		fmt.Println("Cannot Create file!")
		return
	}

	defer file.Close()

	// ---WRITING WITH FPRINTF---
	// fmt.Fprintf(file, "line 1: %s", "new data in the file")

	// ---WRITING THROUGH FILE---
	// n, err := file.WriteString("new data in the file")
	// fmt.Println("no of bytes wrote: ", n)

	// ---WRITING WITH BUFFERED WRITER---
	w := bufio.NewWriterSize(file, 10)
	for i := 0; i < 5; i++ {
		n, _ := w.WriteString("new da")
		fmt.Println("no of bytes wrote: ", n)
		time.Sleep(time.Second * 2)
		w.Flush()
	}

	w.Flush()
}
