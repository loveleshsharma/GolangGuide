package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// func pingTeamA(playerName string, teamName string, chanTeam chan bool, chanPlayer chan bool, chanGameOver chan bool) {
func pingpong(playerName string, teamName string, chanTeam chan bool, chanPlayer chan bool, chanGameOver chan bool) {
	fmt.Println("player created")
	for {
		msgTeam := <-chanTeam
		if msgTeam == false {
			continue
		}
		fmt.Println(teamName, "team chance")
		for {
			msgPlayer := <-chanPlayer
			if msgPlayer == false {
				continue
			}
			fmt.Printf("Team %s, Player %s\n", teamName, playerName)
			randomTime := randomGenerator()
			if randomTime > 500 {
				chanGameOver <- true
			}
			time.Sleep(time.Millisecond * time.Duration(randomTime))
			chanTeam <- !(<-chanTeam)
			chanPlayer <- !(<-chanPlayer)
			break
		}
	}
	// defer close(chanTeam)
	// defer close(chanPlayer)
	// defer close(chanGameOver)
}

func pingTeamB(playerName string, teamName string, chanTeam chan bool, chanPlayer chan bool, chanGameOver chan bool) {

}

func pongTeamA(playerName string, teamName string, chanTeam chan bool, chanPlayer chan bool, chanGameOver chan bool) {

}

func pongTeamB(playerName string, teamName string, chanTeam chan bool, chanPlayer chan bool, chanGameOver chan bool) {

}

func randomGenerator() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(700)
}

func main() {

	channelTeamA := make(chan bool)
	channelAPlayer1 := make(chan bool)
	channelAPlayer2 := make(chan bool)
	channelTeamB := make(chan bool)
	channelBPlayer1 := make(chan bool)
	channelBPlayer2 := make(chan bool)
	channelGameOver := make(chan bool)
	// channelGameOver <- false

	// go pingTeamA("V", "I", channelTeam1, channelPlayer1, channelGameOver)
	// go pingTeamB("D", "I", channelTeam2, channelPlayer1, channelGameOver)
	// go pongTeamA("S", "A", channelTeam1, channelPlayer2, channelGameOver)
	// go pongTeamB("R", "A", channelTeam2, channelPlayer2, channelGameOver)

	go pingpong("V", "I", channelTeamA, channelAPlayer1, channelGameOver)
	go pingpong("D", "I", channelTeamB, channelBPlayer1, channelGameOver)
	go pingpong("S", "A", channelTeamA, channelAPlayer2, channelGameOver)
	go pingpong("R", "A", channelTeamB, channelBPlayer2, channelGameOver)

	// channelTeamA <- true
	// channelAPlayer1 <- true
	// channelAPlayer2 <- false
	// channelTeamB <- false
	// channelBPlayer1 <- false
	// channelBPlayer2 <- false

	for {
		select {
		case msgTeamA := <-channelTeamA:
			channelTeamA <- !(msgTeamA)
			channelTeamB <- msgTeamA
		case msgTeamB := <-channelTeamB:
			channelTeamB <- !(msgTeamB)
			channelTeamA <- msgTeamB
		case msgGameOver := <-channelGameOver:
			if msgGameOver == true {
				os.Exit(1)
			}
		}
	}
	fmt.Println("exiting...")

}
