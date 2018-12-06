package main

import "fmt"

type Person struct {
	name   string
	gender string
	age    int
}

func (p Person) getName() string {
	return p.name
}

// for initializing Person variable. similar to constructor
func newPerson() Person {

	return Person{}
}

func newPersonWithValues(n string, g string, a int) Person {
	return Person{n, g, a}
}

func (p Person) display() {
	fmt.Println("Name: " + p.name)
	fmt.Println("Gender: " + p.gender)
	fmt.Println("Age: ", p.age)
}
