package main

import (
	"fmt"
	"gotutorial/types"
)

// need to run main.go and Person.go together in order to run this

func main() {
	var p1 types.Person
	p1.Name = "Lovelesh"
	p1.Gender = "Male"
	p1.Age = 25
	fmt.Println(p1)

	// structure literal
	p2 := types.Person{"Gautam", "Male", 20}
	fmt.Println(p2)
	// calling receiver function on that structure
	fmt.Println(p2.GetName())

	fmt.Println("------------")
	p3 := types.NewPersonWithValues("Thoughtworks", "Male", 25)
	p3.Display()

	p4 := []types.Person{{}, {}}

}
