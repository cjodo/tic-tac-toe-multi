package websocket

import (
	"fmt"
	"log"

	"github.com/cjodo/tic-tac-toe-multi/pkg/lobby"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID			string
	Player	string 
	Conn		*websocket.Conn
}

type Message struct {
	Type int		`json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()

		if err != nil {
			log.Println("Client err:", err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}

		fmt.Printf("Move Received: %v from: %v\n", message, c.ID)
	}
}
