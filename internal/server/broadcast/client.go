package broadcast

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	hub     *Hub
	conn    *websocket.Conn
	send    chan Message
	channel string
}

func ServeWs(hub *Hub, c *gin.Context, channel string) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &Client{hub: hub, conn: conn, send: make(chan Message), channel: channel}
	client.hub.register <- client

	go client.readPump()
	go client.writePump()
}

func (c *Client) writePump() {
	defer func() {
		err := c.conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				err := c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					fmt.Println(err)
				}
				return
			}

			err := c.conn.WriteMessage(websocket.TextMessage, message.data)
			if err != nil {
				return
			}
		}
	}
}

func (c *Client) readPump() {
	defer func() {
		err := c.conn.Close()
		c.hub.unregister <- c
		if err != nil {
			fmt.Println(err)
		}
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		c.hub.broadcast <- Message{c, message, c.channel}
	}
}
