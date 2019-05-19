package main

import (
	"fmt"
	"math/rand"
	"time"
)

func cakeMaker(kind string, number int, to chan<- string) {
	rand.Seed(time.Now().Unix())
	for i := 0; i <  number; i++ {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		to <- kind
	}
	close(to)
}

func main() {
	chocolateChan := make(chan string, 8)
	redVelVetChan := make(chan string, 8)

	go cakeMaker("chocolate",4,chocolateChan)
	go cakeMaker("redvelvet",3,redVelVetChan)

	moreChocolate := true
	moreRedVelVet := true

	var cake string

	for moreChocolate || moreRedVelVet {
		select {
			case cake, moreChocolate = <-chocolateChan :
				if moreChocolate == true {
					fmt.Printf("Delivered %s\n",cake)
				}
			case cake, moreRedVelVet = <-redVelVetChan :
				if moreRedVelVet == true {
					fmt.Printf("Delivered %s\n",cake)
				}
		}


	}
}
