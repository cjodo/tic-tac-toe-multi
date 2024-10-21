package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
	"sync"

	"github.com/cjodo/tic-tac-toe-multi/pkg/message"
	"github.com/cjodo/tic-tac-toe-multi/pkg/websocket"
)

const (
	MESSAGE_TYPE_NOTIFY = 1
	MESSAGE_TYPE_MOVE = 2
)

type Lobby struct {
	Broadcast				chan interface{}
	Register				chan *websocket.Client
	Unregister			chan *websocket.Client
	WaitingPlayers 	[]*websocket.Client
	ActiveGames			map[string]*Game
	Clients					map[*websocket.Client]bool
	mu						sync.Mutex
}

func NewLobby() *Lobby {
	return &Lobby{
		WaitingPlayers: []*websocket.Client{},
		ActiveGames:		make(map[string]*Game),
		Broadcast:			make(chan interface{}),
		Register:				make(chan *websocket.Client),
		Unregister:			make(chan *websocket.Client),
		Clients:				make(map[*websocket.Client]bool),
	}
}

func (l *Lobby) Run() {
	for {
		select {
		case client := <-l.Register:
			l.mu.Lock()
			l.WaitingPlayers = append(l.WaitingPlayers, client)
			l.PairPlayers()
			l.mu.Unlock()

		case client := <-l.Unregister:
			l.mu.Lock()
			l.deleteClient(client)
			l.mu.Unlock()

		case msg := <-l.Broadcast:
			l.mu.Lock()
			fmt.Println("Move in broadcast: ", msg)
			// go handleMessage(msg);

			l.mu.Unlock()
		}
	}
}

func (l *Lobby) HandleNewConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil{
		fmt.Println("WebSocket Upgrade Failed: ", err)
		return
	}

	random := rand.IntN(100);

	client := &websocket.Client{
		ID:			strconv.Itoa(random),
		Conn:		conn,
		Send:		make(chan interface{}),
	}

	if _, ok := l.Clients[client]; !ok {
		l.Clients[client] = true

		l.Register <- client
		go l.handleClientRead(client)
		go l.handleClientWrite(client)
	}
}

func (l *Lobby) handleClientRead(client *websocket.Client) {
	defer func() {
		l.Unregister <- client
		client.Conn.Close()
	}()

	for {
		var rawMessage json.RawMessage 
		err := client.Conn.ReadJSON(&rawMessage)

		if err != nil {
			fmt.Printf("Error reading from client %s: %v:", client.ID, err)
			break
		}

		var msgMap map[string]interface{}

		if err := json.Unmarshal(rawMessage, &msgMap); err != nil {
			fmt.Printf("Error unmarshalling raw message: %v:", err)
			break
		}

		if err := l.handleMessage(rawMessage, msgMap); err != nil {
			fmt.Printf("Error handling message: %v:", err)
			break
		}
	}
}

func (l *Lobby) handleClientWrite(client *websocket.Client) {
	defer func() {
		l.Unregister <- client
		client.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-client.Send:
			fmt.Println("Recieved message: ", msg)

			if !ok {
				fmt.Println("Error Sending message")
				client.Conn.WriteMessage(2, []byte{})
				return
			}

			err := client.Conn.WriteJSON(msg)
			if err != nil {
				fmt.Printf("Error writing to client %s: %v", client.ID, err)
				return
			}
		}
	}
}

func (l *Lobby) PairPlayers() {
	if(len(l.WaitingPlayers) < 2) {
		return
	}

	player1 := l.WaitingPlayers[0]
	player2 := l.WaitingPlayers[1]

	player1.Player = "X"
	player2.Player = "O"

	fmt.Printf("Pairing Players %v, %v\n", player1.ID, player2.ID)

	game := NewGame(player1, player2)

	game.GameID = "TestGame"

	l.ActiveGames["TestGame"] = game

	startPlayer1 := message.NewStartGame("X", game.GameID, player1.ID)
	startPlayer2 := message.NewStartGame("O", game.GameID, player2.ID)

	player1.Send <- startPlayer1
	player2.Send <- startPlayer2
}

func (l *Lobby) handleMessage(raw json.RawMessage, msgMap map[string]interface{}) error {
	if _, ok := msgMap["type"]; !ok {
		return errors.New("Type not found in message")
	}

	msgType := msgMap["type"]

	switch msgType {
	case "move":
	var msg message.MoveMessage
		if err := json.Unmarshal(raw, &msg); err != nil {
			return errors.New("Error unmarshalling rawMessage")
		}

		from := l.findClientById(msg.PlayerId)
		game := l.ActiveGames[msg.GameId]

		game.RegisterMove(msg)

		opponent := l.findOpponent(game, from)
		if opponent == nil {
			return errors.New("opponent not found")
		}
		// Try clearing the chan
		opponent.Send <- msg
	}

	return nil
}

func (l *Lobby) findOpponent(game *Game, from *websocket.Client) *websocket.Client {
	for player, _ := range game.Players {
		if player != from {
			return player
		}
	}
	return nil
}

func (l *Lobby) findClientById(id string) *websocket.Client {
	for client, _ := range l.Clients {
		if client.ID == id {
			return client
		}
	}
	return nil
}

func (l *Lobby) deleteClient(client *websocket.Client) {
	fmt.Println("Deleting client: ", client.ID)

	delete(l.Clients, client);
	if _, ok := <-client.Send; ok {
		close(client.Send)
	}

	for i, player := range l.WaitingPlayers {
		if(player.ID == client.ID) {
			l.WaitingPlayers[i] = l.WaitingPlayers[len(l.WaitingPlayers) - 1]
			l.WaitingPlayers = l.WaitingPlayers[:len(l.WaitingPlayers) - 1]
		}
	} 

	//Remove from active game
	for _, game := range l.ActiveGames {
		delete(game.Players, client)
	} 
}
