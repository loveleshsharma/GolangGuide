package main

import (
	"fmt"
	"sync"
)

var (
	value int
)

func produce(val *int, ch chan bool, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		// *val = *val + 1
		value = value + 1
		// ch <- true
	}
	wg.Done()
}

func consume(val *int, ch chan bool, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		// <-ch
		// *val = *val - 1
		value = value - 1
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	val := 100
	ch := make(chan bool)

	wg.Add(2)
	go produce(&val, ch, &wg)
	go consume(&val, ch, &wg)

	wg.Wait()
	fmt.Println("Expected value = 100 and Actual value = ", val)
}
