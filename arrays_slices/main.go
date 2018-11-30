package main

import (
	"fmt"
)

func main() {

	var arr [2]string
	arr[0] = "One"
	arr[1] = "Two"

	// arr := [2]string{"One", "Two"}

	// Slice with dynamic size
	arrSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(arr)
	fmt.Println(arrSlice)

	a := [5]int{1, 2, 3, 4, 5}

	// cannot pass array into slice parameter
	// sliceExp(a)

	// Slice initialized from array a with the range
	aSlice := a[0:2]

	// the changes made in sliceExp function reflects back here
	sliceExp(aSlice)

	fmt.Println(aSlice)

	fmt.Println("--------------------")

	// another way to create a slice
	newSlice := make([]int, 5, 10)
	newSlice = append(newSlice, 10)
	newSlice = append(newSlice, 20)
	newSlice = append(newSlice, 30)
	newSlice = append(newSlice, 40)
	newSlice = append(newSlice, 50)
	newSlice1 := append(newSlice, 60)

	fmt.Println("newSlice1: ", newSlice1)

	// newSlice := []int{10, 20, 30, 40, 50}
	fmt.Println("newSlice: ", newSlice)

	// setting slice length to zero
	newSlice = newSlice[:0]
	fmt.Println(newSlice)

	// setting slice length to 3
	newSlice = newSlice[:3]
	fmt.Println(newSlice)

	fmt.Println(newSlice[1])

	fmt.Println("***********************")

	slice1 := make([]int, 5, 5)

	for i := 0; i < 5; i++ {
		slice1[i] = (i + 1) * 10
	}

	fmt.Println("slice1: ", slice1)
	fmt.Println("len: ", len(slice1))
	fmt.Println("cap: ", cap(slice1))

	slice2 := slice1[0:2]
	// changes made to slice2 will reflect to slice1 also, because it is a reference type
	slice2[0] = 99

	fmt.Println(slice2)
	fmt.Println(slice1)

	slice3 := append(slice1, 1, 2, 3, 4, 5, 6, 7)

	fmt.Println("slice1: ", slice1)
	fmt.Println("slice3: ", slice3)
	fmt.Println("len: ", len(slice1))
	fmt.Println("cap: ", cap(slice1))

	fmt.Println("######################")
	sliceA := []int{100, 200, 300, 400, 500}
	sliceB := make([]int, len(sliceA))
	// copying one slice into another
	copy(sliceB, sliceA)
	fmt.Println("sliceB: ", sliceB)

	sliceC := sliceA[2:]
	sliceC[0] = 99
	fmt.Println("SliceA: ", sliceA)

	// copying range of a slice into another
	sliceA = []int{100, 200, 300, 400, 500}
	sliceB = make([]int, len(sliceA))

	copy(sliceB[0:2], sliceA[3:])

	fmt.Println("sliceB: ", sliceB)

}

func sliceExp(arr []int) {

	arr[0] = 99

}
