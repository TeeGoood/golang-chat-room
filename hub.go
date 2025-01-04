package main

type hub struct {
	clients map[*client]bool
}

func newHub() *hub {
	return &hub{
		clients: make(map[*client]bool),
	}
}

func (h *hub) boardcast(msg string) {
	for client := range h.clients {
		client.send(msg)
	}
}

func (h *hub) addClient(c *client) {
	h.clients[c] = true
}

func (h *hub) removeClient(c *client) {
	delete(h.clients, c)
}
