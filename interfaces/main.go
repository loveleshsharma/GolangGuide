package main

import "fmt"

type animal interface {
	speak() string
}

type cat struct{}
type dog struct{}

func main() {
	c := cat{}
	d := dog{}
	animalVoice(c)
	animalVoice(d)
}

func (c cat) speak() string {
	return "Meow"
}

func (d dog) speak() string {
	return "Bow Wow"
}

func (d dog) say() string {
	return "Bowwwwooooo"
}

func animalVoice(a animal) {
	fmt.Println(a.speak())
}
