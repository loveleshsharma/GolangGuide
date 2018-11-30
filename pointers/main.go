package main

import (
	"fmt"
)

func main() {
	a := 10

	// declaring pointer variable
	var b *int

	// assigning address of a into b
	b = &a

	// printing value of memory address pointed by b
	fmt.Println(*b)

	// changing the value at the mem address pointed by b
	*b = 20

	// changes are reflected to the original variable a
	fmt.Println(a)

}
