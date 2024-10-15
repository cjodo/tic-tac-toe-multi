package lobby

import (
	"fmt"
	"sync"

	"github.com/cjodo/tic-tac-toe-multi/pkg/websocket"
)

type Lobby struct {
		WaitingPlayers 	[]*websocket.Client
		Mutex						sync.Mutex
}

func NewLobby() *Lobby {
		return &Lobby{
				WaitingPlayers: []*websocket.Client{},
		}
}

func (l *Lobby) AddPlayer(c *websocket.Client) {
		l.Mutex.Lock()
		defer l.Mutex.Unlock()

		l.WaitingPlayers = append(l.WaitingPlayers, c)

		fmt.Println("Added Player: ", c.ID)

		if(len(l.WaitingPlayers) >= 2) {
				l.PairPlayers()
		}
}

func (l *Lobby) PairPlayers() {
		return
}
