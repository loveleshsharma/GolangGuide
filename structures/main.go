package main

import "fmt"

// need to run main.go and Person.go together in order to run this

func main() {
	var p1 Person
	p1.name = "Lovelesh"
	p1.gender = "Male"
	p1.age = 25
	fmt.Println(p1)

	// structure literal
	p2 := Person{"Gautam", "Male", 20}
	fmt.Println(p2)
	// calling receiver function on that structure
	fmt.Println(p2.getName())

	fmt.Println("------------")
	p3 := newPersonWithValues("Thoughtworks", "Male", 25)
	p3.display()

}
