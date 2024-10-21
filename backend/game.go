package main

import (
	"fmt"

	"github.com/cjodo/tic-tac-toe-multi/pkg/message"
	"github.com/cjodo/tic-tac-toe-multi/pkg/websocket"
)

type Game struct {
	GameID		string
	Board			[9]string
	Players		map[*websocket.Client]bool
}

type Move struct {
	Game		*Game
	Sender	*websocket.Client
	Move		int
}

func NewGame(player1, player2 *websocket.Client) *Game {
	game := &Game{
		Board:		[9]string{},
		Players:	make(map[*websocket.Client]bool),
	}

	game.Players[player1] = true
	game.Players[player2] = true

	return game
}

func (g *Game) checkStatus() {
	//Implement win lose draw checking
}

func (g *Game) RegisterMove(msg message.MoveMessage) {
	g.Board[msg.Move] = msg.Player
	g.printBoard()
}

func (g *Game) printBoard() {
	for i := range g.Board {
			fmt.Printf("[%v]", g.Board[i])
	}
}
