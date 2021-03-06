package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	gameChan = make(chan bool)
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
	} else {
		*t.currentPlayer = t.player1
	}
}

func (t team) teamPlay() {
	t.togglePlayer()
	t.currentPlayer.channel <- true	//sending the play signal to the respective player
}

func (t team) getCurrentPlayer() player {
	return *t.currentPlayer
}

func (t team) startGame() {
	go t.player1.play()
	go t.player2.play()
}

func newPlayer(playerName string) player {
	plr := player{}
	plr.name = playerName
	plr.channel = make(chan bool)
	return plr
}

func (p player) play() {
	fmt.Println(p.name ,"is ready!")
	for{
		<-p.channel
		millis := randomGenerator()
		if millis <= 500 {
			fmt.Printf("%s          : played in %dms\n",p.name,millis)
		} else {
			fmt.Println("Game is finished....")
			fmt.Printf("%s          : took more than 500ms to play his shot, and hence their team lost the match.\n",p.name)
			gameChan<-false
		}
		gameChan<-true
	}
}

func randomGenerator() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(700)
}

func main() {

	plr1 := newPlayer("Virat")
	plr2 := newPlayer("Dhoni")
	plr3 := newPlayer("Steve")
	plr4 := newPlayer("Ricky")

	tem1 := newTeam("India", plr1, plr2)
	tem2 := newTeam("Australia", plr3, plr4)
	var currTem team

	//manually initiating the Game
	go func() {
		fmt.Println("Starting a new Game...")
		tem1.startGame()
		tem2.startGame()
		gameChan <- true
	}()

	for {
		res := <-gameChan
		if res == true {

			time.Sleep(time.Millisecond*700)
			if currTem == tem1 {
				currTem = tem2
			} else if currTem == tem2 {
				currTem = tem1
			} else {
				currTem = tem1
			}

			currTem.teamPlay()

		} else {
			fmt.Println("exiting...")
			break
		}
	}
}
