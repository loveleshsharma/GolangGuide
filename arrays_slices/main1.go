package main

import "fmt"

func main() {

	var myslice []string
	var size = 100

	for i := 0; i < 2; i++ {

		if myslice == nil {
			fmt.Println("nil")
			myslice = make([]string, size)
		} else {
			fmt.Println("not nil")
		}
	}

}
