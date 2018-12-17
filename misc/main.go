package main

import (
	"fmt"
	"reflect"
)

type myinterface interface {
}

func main() {

	s := []int{1, 2, 3, 4, 5}
	// s = append(s, 6)

	// delete the 3rd element
	var i = 3
	s = append(s[:i], s[i+1:]...)

	fmt.Println(s)

	fmt.Println("Hello World")

	myfunc("hello")
	myfunc(1)
	myfunc(1.0)
	myfunc(byte(12))

}

func myfunc(i interface{}) {
	fmt.Println(reflect.TypeOf(i).String())
}
