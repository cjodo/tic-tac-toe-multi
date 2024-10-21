package websocket

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	ID			string
	Conn		*websocket.Conn
	Player	string
	Send		chan interface{}
}

func NewClient(Conn *websocket.Conn, ID, Player string) *Client {
	return &Client{
		ID:			ID,
		Conn:		Conn,
		Player: Player,
		// Allows multiple data types to be sent ** see pkg/message
		Send:		make(chan interface{}),
	}
}

