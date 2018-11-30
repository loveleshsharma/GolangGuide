package main

import (
	"fmt"
)

func main() {
	var color string = "blue"

	// if else if else
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("it is neither red nor blue")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("it is neither red nor blue")
	}

}
