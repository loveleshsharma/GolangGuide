package main

import "fmt"

func main() {

	// defining map
	names := make(map[int]string)

	// now assigning key and values
	names[100] = "Lovelesh"
	names[101] = "Thoughtworks"
	names[102] = "OnePlus"

	fmt.Println(names)
	fmt.Println(names[100])
	fmt.Println(len(names))

	// deleting from map
	delete(names, 101)
	fmt.Println(names)

	// defining and initializing map
	games := map[int]string{1: "mario", 2: "contra", 3: "pubG", 4: "subway surfers"}

	fmt.Println(games)

}
