package main

import (
	"fmt"
	"log"
	"net/http"
)

var PORT = ":4040"

func setupRoutes() {
	lobby := NewLobby()

	go lobby.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		lobby.HandleNewConnection(w, r)
	})
}

func main() {
	fmt.Printf("Server running on port: %v\n", PORT)

	setupRoutes()

	log.Fatal(http.ListenAndServe(":4040", nil))
}
