package main

import (
	"log"
	"net/http"
)

func main() {
	hub := newHub()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsHandler(hub, w, r)
	})
	addr := "0.0.0.0:8000"
	log.Printf("listen to %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
