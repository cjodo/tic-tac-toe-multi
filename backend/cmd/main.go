package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/cjodo/tic-tac-toe-multi/pkg/websocket"
)

var PORT = ":4040"

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}

	rand := rand.IntN(200)
	
	client := &websocket.Client{
		ID: strconv.Itoa(rand),
		Conn: ws,
	}

	fmt.Printf("Client joined %v\n", client.ID)

	client.Read()
}

func setupRoutes() {

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r)
	})
}

func main() {
	fmt.Printf("Server running on port: %v\n", PORT)
	setupRoutes()

	log.Fatal(http.ListenAndServe(":4040", nil))
}
