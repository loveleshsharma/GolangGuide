package main

import (
	"fmt"
	"sync"
	"time"
)

func incommingChannel(ch <-chan string, wg *sync.WaitGroup) {
	fmt.Println("Waiting for message...")
	fmt.Println(<-ch)
	wg.Done()
}

func main() {

	ch := make(chan string, 1)
	var wg sync.WaitGroup
	// for i := 0; i < 10; i++ {
	wg.Add(1)
	go incommingChannel(ch, &wg)
	// }

	time.Sleep(time.Second)

	ch <- "First Message"
	// wg.Wait()
}
