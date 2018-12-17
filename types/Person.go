package types

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name   string
	Gender string
	Age    int
}

func (p Person) GetName() string {
	return p.Name
}

// for initializing Person variable. similar to constructor
func NewPerson() Person {

	return Person{}
}

func NewPersonWithValues(n string, g string, a int) Person {
	return Person{n, g, a}
}

func (p Person) Display() {
	fmt.Println("Name: " + p.Name)
	fmt.Println("Gender: " + p.Gender)
	fmt.Println("Age: ", p.Age)
	fmt.Printf("%+v\n", p)
}

// ToString converts the Person object into a single string
func (p Person) ToString() string {

	return p.Name + "," + p.Gender + "," + strconv.Itoa(p.Age)
}

// SaveToFile saves the Person object into a File on the HardDrive
func (p Person) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(p.ToString()), 0666)
}

func ReadFromFile(filename string) (Person, error) {
	bytes, err := ioutil.ReadFile(filename)
	var person Person
	if err != nil {
		fmt.Println("Problem reading file! Exiting...")
		os.Exit(1)
	} else {
		data := string(bytes)
		stringArray := strings.Split(data, ",")
		val, _ := strconv.Atoi(strings.Trim(stringArray[2], " "))
		person = NewPersonWithValues(stringArray[0], stringArray[1], val)
	}
	return person, err
}
