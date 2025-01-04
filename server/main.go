package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Handler struct{}

func (h *Handler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world\n"))
}

func (h *Handler) StartWebSocket(w http.ResponseWriter, r *http.Request) {
	log.Printf("url: %s", r.URL)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error start websocket: %v", err)
	}
	conn.WriteMessage(websocket.TextMessage, []byte("start websocket success"))
	err = conn.Close()
	if err != nil {
		log.Printf("erorr close websocket: %s", err)
	}
}

func main() {
	handler := Handler{}
	http.HandleFunc("/", handler.HelloWorld)
	http.HandleFunc("/ws", handler.StartWebSocket)

	addr := "0.0.0.0:8000"
	log.Printf("listen to %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
