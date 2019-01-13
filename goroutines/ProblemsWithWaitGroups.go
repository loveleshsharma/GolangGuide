package main

import (
	"fmt"
	"sync"
)

func increment(ptr *int, wg *sync.WaitGroup) {
	i := *ptr
	fmt.Println("i=", i)
	*ptr = i + 1
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var value = 0

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go increment(&value, &wg)
	}

	wg.Wait()
	fmt.Println("final value", value)

}
