package broadcast

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	Stream  *Stream
	Conn    *websocket.Conn
	Send    chan Message
	Channel string
    Type    string
}

func ServeWs(hub *Hub, c *gin.Context, channel string, clientType string) (*Client, error) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
    stream, ok := hub.Streams[channel]

    if !ok {
        stream = Stream{Hub: hub, Ip: "", Clients: make([]*Client, 2)}
        hub.Streams[channel] = stream
    }

    client := &Client{Stream: &stream, Conn: conn, Send: make(chan Message), Channel: channel, Type: clientType}
	client.Stream.Hub.Register <- client

	go client.readPump()
	go client.writePump()

    return client, nil
}

func (c *Client) writePump() {
	defer func() {
		c.Stream.Hub.Unregister <- c
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					fmt.Println(err)
				}
				return
			}

			err := c.Conn.WriteMessage(websocket.TextMessage, message.Data)
			if err != nil {
				return
			}
		}
	}
}

func (c *Client) readPump() {
	defer func() {
		c.Stream.Hub.Unregister <- c
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		c.Stream.Hub.Broadcast <- Message{c, message, c.Channel}
	}
}
