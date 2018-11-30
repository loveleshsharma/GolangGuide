package main

import "fmt"

func main() {

	fmt.Println("hello world")
	fmt.Println(Greeting("Lovelesh"))
	fmt.Println(Sum(55, 5))

}

func Greeting(name string) string {
	return "Hello " + name
}

func Sum(a int, b int) int {
	return a + b
}
