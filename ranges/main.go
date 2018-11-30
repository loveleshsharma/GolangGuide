package main

import (
	"fmt"
)

func main() {
	ids := []int{10, 20, 30, 40, 50}

	// RANGE WITH SLICE

	// loop though ids with index
	for i, id := range ids {
		fmt.Println(i, " - ID: ", id)
	}

	// without using index
	for _, id := range ids {
		fmt.Println("ID: ", id)
	}

	fmt.Println("-----Maps here-----")
	// RANGE WITH MAP

	// defining and initializing map
	games := map[int]string{1: "mario", 2: "contra", 3: "pubG", 4: "subway surfers"}

	// looping with keys and values
	for k, v := range games {
		fmt.Printf("KEY: %d, VALUE: %s\n", k, v)
	}

	// if you want to ignore values
	fmt.Println("KEYS:")
	for k := range games {
		fmt.Println(k)
	}

	// if you want to ignore keys
	fmt.Println("VALUES:")
	for _, v := range games {
		fmt.Println(v)
	}
}
