package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wsHandler(hub *hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrate connection: %s", err)
		return
	}

	client := newClient(conn, hub)
	client.read()
}

func main() {
	hub := newHub()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsHandler(hub, w, r)
	})
	addr := "0.0.0.0:8000"
	log.Printf("listen to %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
