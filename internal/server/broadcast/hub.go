package broadcast

import (
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Hub struct {
	Clients    map[string]map[*Client]bool
	Broadcast  chan Message
	Register   chan *Client
	Unregister chan *Client
}

type Message struct {
	Sender  *Client
	Data    []byte
	Channel string
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan Message),
		Clients:    make(map[string]map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			_, ok := h.Clients[client.Channel]
			if !ok {
				h.Clients[client.Channel] = make(map[*Client]bool)
			}
			h.Clients[client.Channel][client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.Channel][client]; ok {
				delete(h.Clients[client.Channel], client)
			}
		}
	}
}
