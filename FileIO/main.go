package main

import (
	"fmt"
	"gotutorial/types"
)

func main() {

	person := types.NewPerson()
	person.Name = "Jack"
	person.Gender = "Male"
	person.Age = 35

	person.Display()

	byteSlice := []byte(person.ToString())

	fmt.Println(byteSlice)
	person.SaveToFile("my_person")

	myPerson, _ := types.ReadFromFile("my_person")
	fmt.Println("From HardDrive: ", myPerson)
}
