package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cjodo/pkg/websocket"
)


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the room")
}


func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	
	reader(ws)
}


func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App")
	setupRoutes()

	log.Fatal(http.ListenAndServe(":4040", nil))
}
