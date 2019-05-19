package main

import (
	"fmt"
	"sync"
	"time"
)

type safeStruct struct {
	counter int
	lock    sync.Mutex
}

func (ss *safeStruct) increment() {
	ss.lock.Lock()
	ss.counter++
	ss.lock.Unlock()
}

func main() {

	structVar := safeStruct{}
	for i := 0; i < 1000; i++ {
		go structVar.increment()
	}

	time.Sleep(time.Second)
	fmt.Println(structVar.counter)

}
