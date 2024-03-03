package broadcast

import (
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Hub struct {
	Streams    map[string]*Stream
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
		Streams:    make(map[string]*Stream),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			stream, ok := h.Streams[client.Channel]
			if !ok {
                stream = &Stream{Ip: "", Clients: make([]*Client, 2)}
                h.Streams[client.Channel] = stream
			}

			stream.Clients = append(stream.Clients, client)

		case client := <-h.Unregister:
            stream, ok := h.Streams[client.Channel]
            if !ok {
                break
            }

            for i, c := range h.Streams[client.Channel].Clients {
                if client == c {
                    stream.Clients = append(stream.Clients[:i], stream.Clients[i+1:]...)
                    break
                }
            }
		}
	}
}
