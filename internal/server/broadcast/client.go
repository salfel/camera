package broadcast

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	Hub     *Hub
	Conn    *websocket.Conn
	Send    chan Message
	Channel string
}

func ServeWs(hub *Hub, c *gin.Context, channel string) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &Client{Hub: hub, Conn: conn, Send: make(chan Message), Channel: channel}
	client.Hub.Register <- client

	go client.readPump()
	go client.writePump()
}

func (c *Client) writePump() {
	defer func() {
		c.Hub.Unregister <- c
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
		c.Hub.Unregister <- c
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		c.Hub.Broadcast <- Message{c, message, c.Channel}
	}
}
