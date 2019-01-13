package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	tAch1 = make(chan bool)
	tAch2 = make(chan bool)
	tBch1 = make(chan bool)
	tBch2 = make(chan bool)
	GO    = make(chan string)
	teamA = "India"
	teamB = "Australia"
)

func tAP1(tName string, pName string, player chan bool, GO chan string) {
	fmt.Println(pName, " - Ready")
	for {
		<-player
		time.Sleep(time.Millisecond * 300)
		randomTime := randomGenerator()
		if randomTime <= 500 {
			fmt.Printf("%s          : played in %dms\n", pName, randomTime)
		} else {
			fmt.Printf("%s          : took more than 500ms to play his shot, and hence their team lost the match.\n", pName)
			GO <- fmt.Sprintf("Team %s has WON", teamB)
		}
		tBch1 <- true
	}

}

func tAP2(tName string, pName string, player chan bool, GO chan string) {
	fmt.Println(pName, " - Ready")
	for {
		<-player
		time.Sleep(time.Millisecond * 300)
		randomTime := randomGenerator()
		if randomTime <= 500 {
			fmt.Printf("%s          : played in %dms\n", pName, randomTime)
		} else {
			fmt.Printf("%s          : took more than 500ms to play his shot, and hence their team lost the match.\n", pName)
			GO <- fmt.Sprintf("Team %s has WON", teamB)
		}
		tBch2 <- true
	}

}

func tBP1(tName string, pName string, player chan bool, GO chan string) {
	fmt.Println(pName, " - Ready")
	for {
		<-player
		time.Sleep(time.Millisecond * 300)
		randomTime := randomGenerator()
		if randomTime <= 500 {
			fmt.Printf("%s          : played in %dms\n", pName, randomTime)
		} else {
			fmt.Printf("%s          : took more than 500ms to play his shot, and hence their team lost the match.\n", pName)
			GO <- fmt.Sprintf("Team %s has WON", teamA)
		}
		tAch2 <- true

	}
}

func tBP2(tName string, pName string, player chan bool, GO chan string) {
	fmt.Println(pName, " - Ready")
	for {
		<-player
		time.Sleep(time.Millisecond * 300)
		randomTime := randomGenerator()
		if randomTime <= 500 {
			fmt.Printf("%s          : played in %dms\n", pName, randomTime)
		} else {
			fmt.Printf("%s          : took more than 500ms to play his shot, and hence their team lost the match.\n", pName)
			GO <- fmt.Sprintf("Team %s has WON", teamA)
		}
		tAch1 <- true
	}
}

func randomGenerator() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(700)
}

func main() {
	go func() {
		fmt.Println("Starting a new Game")
		time.Sleep(time.Second)
		tAch1 <- true
	}()

	go tAP1(teamA, "Virat", tAch1, GO)
	go tAP2(teamA, "Dhoni", tAch2, GO)
	go tBP1(teamB, "Steve", tBch1, GO)
	go tBP2(teamB, "Ricky", tBch2, GO)

	for {
		select {
		case msg := <-GO:
			fmt.Println("Game is finished....")
			fmt.Println(msg)
			return
		}
	}
}
