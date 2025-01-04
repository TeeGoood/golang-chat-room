package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var message = make(chan string)
var clients = make(map[chan string]bool)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func hub() {
	for {
		m := <-message
		for client := range clients {
			client <- m
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("failed create websocket connection: %s", err)
		return
	}
	get := make(chan string)
	clients[get] = true
	done := make(chan int)

	// write
	go func() {
		defer func() {
			done <- 1
		}()
		for {
			m := <-get
			if err := conn.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
				log.Printf("failed write message: %s", err)
				return
			}
		}
	}()

	// read
	go func() {
		defer func() {
			done <- 1
		}()
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Printf("failed read message: %s", err)
				return
			}

			if messageType == websocket.TextMessage {
				log.Println(string(p))
				message <- string(p)
			}
		}
	}()

	// disconnect
	go func() {
		<-done
		log.Println("close socket")
		if err := conn.Close(); err != nil {
			log.Printf("error close socket: %s", err)
		}
		delete(clients, get)
	}()
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsHandler)
}

func main() {
	setupRoutes()
	go hub()
	addr := "0.0.0.0:8000"
	log.Printf("listen to %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
