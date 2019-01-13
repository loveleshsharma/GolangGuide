package main

import (
	"fmt"
	"sync"
)

func receivefrombuffer(ch <-chan int, wg *sync.WaitGroup) {
	fmt.Println("RECEIVER: Waiting for the buffer to be full...")
	fmt.Println(<-ch)
	fmt.Println("RECEIVER: going to exit")
	wg.Done()
}

func sendtobuffer(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("SENDER: sending to the buffer...")
	for i := 0; i < 3; i++ {
		ch <- 1
	}
}

func main() {

	ch := make(chan int, 3)
	var wg sync.WaitGroup
	wg.Add(1)
	go receivefrombuffer(ch, &wg)
	sendtobuffer(ch, &wg)
	wg.Wait()
}
