package broadcast

import (
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Hub struct {
	clients    map[string]map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

type Message struct {
	sender  *Client
	data    []byte
	channel string
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		clients:    make(map[string]map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			_, ok := h.clients[client.channel]
			if !ok {
				h.clients[client.channel] = make(map[*Client]bool)
			}
			h.clients[client.channel][client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client.channel][client]; ok {
				delete(h.clients[client.channel], client)
			}
			close(client.send)
		case message := <-h.broadcast:
			for client := range h.clients[message.channel] {
				if message.channel != client.channel {
					continue
				}

				client.send <- message
			}
		}
	}
}
