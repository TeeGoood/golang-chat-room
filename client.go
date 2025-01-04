package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type client struct {
	conn *websocket.Conn
	hub  *hub
}

func newClient(conn *websocket.Conn, hub *hub) *client {
	client := client{conn: conn, hub: hub}
	hub.addClient(&client)
	return &client
}

func (c *client) send(msg string) {
	if err := c.conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		c.disconnect()
	}
}

func (c *client) read() {
	defer c.disconnect()
	for {
		msgType, p, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("error read message: %s", err)
			break
		}

		if msgType == websocket.TextMessage {
			c.hub.boardcast(string(p))
		}
	}
}

func (c *client) disconnect() {
	c.hub.removeClient(c)
	if err := c.conn.Close(); err != nil {
		log.Printf("error close websocket: %s", err)
	}
}
