package main

import (
	"fmt"
	"math/rand"
	"time"
)

type player struct {
	name    string
	channel chan bool
}

type team struct {
	player1       player
	player2       player
	currentPlayer *player
	name          string
}

func newTeam(teamName string, player1 player, player2 player) team {
	tem := team{}
	tem.player1 = player1
	tem.player2 = player2
	tem.name = teamName
	tem.currentPlayer = &tem.player1
	return tem
}

func (t team) togglePlayer() {
	if *t.currentPlayer == t.player1 {
		*t.currentPlayer = t.player2
	} else if *t.currentPlayer == t.player2 {
		*t.currentPlayer = t.player1
	}
}

func (t team) getCurrentPlayer() player {
	return *t.currentPlayer
}

func newPlayer(playerName string) player {
	plr := player{}
	plr.name = playerName
	plr.channel = make(chan bool)
	return plr
}

func (p player) play() int {
	return randomGenerator()
}

func randomGenerator() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(700)
}

func main() {

	gameChan := make(chan int)
	plr1 := newPlayer("Virat")
	plr2 := newPlayer("Dhoni")
	plr3 := newPlayer("Steve")
	plr4 := newPlayer("Ricky")

	tem1 := newTeam("India", plr1, plr2)
	tem2 := newTeam("Australia", plr1, plr2)

}

func pingpong(plr player, gameChan chan string) {

	for {
		<-gameChan
		millis := plr.play()
		if millis <= 500 {
			fmt.Printf("%s          : played in %dms\n",plr.name,millis)
		} else {
			fmt.Printf("%s          : took more than 500ms to play his shot, and hence their team lost the match.\n",plr.name)

		}


	}

}
