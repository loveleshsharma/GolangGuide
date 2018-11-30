package main

import (
	"fmt"
	"gotutorial/packages"
	"math/rand"
)

func main() {
	fmt.Println("the quick brown fox jumps over the lazy dog!")
	fmt.Println("\nRandom Number:", rand.Intn(10))

	// shorthand variable declaration
	// name := "macbook"
	// var age = 42
	const isCool = true
	// isCool = false
	var floatexp = 2.5

	var name, age = "Ronaldo", 5

	fmt.Printf("hey: %s", name)
	fmt.Printf("\nhello")
	fmt.Printf("\nmy age %d", age)
	fmt.Printf("\nisCool: %T\n", floatexp)
	fmt.Println(name, age, isCool, floatexp)

	fmt.Println(packages.Add(5, 6))
	fmt.Println("hello " + name)

	// fmt.Println(functions.Greeting("Lovelesh"))
}
