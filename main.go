package main

import (
	"log"
	"net/http"

	"minesweeper/game"

	"golang.org/x/net/websocket"
)

func main() {
	server := game.NewServer()
	http.Handle("/ws", websocket.Handler(server.Wshandler))
	http.Handle("/", http.FileServer(http.Dir("./resources")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
