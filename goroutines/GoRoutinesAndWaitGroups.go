package main

import (
	"fmt"
	"sync"
)

func myFunc(i int, wg *sync.WaitGroup) {
	fmt.Println("i=", i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go myFunc(i, &wg)
	}

	wg.Wait()
}
