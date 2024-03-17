package broadcast

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second

	poingWait  = 60 * time.Second
	pingPeriod = (poingWait * 9) / 10

	maxMessageSize = 512
)

type Client struct {
	Stream  *Stream
	Conn    *websocket.Conn
	Send    chan Message
	Message chan Message
	Channel string
	Type    string
}

type MessageType struct {
	Type string `json:"type"`
}

func ServeWs(hub *Hub, c *gin.Context, channel string, clientType string) (*Client, context.Context, error) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	stream, ok := hub.Streams[channel]

	if !ok {
		stream = &Stream{Hub: hub, Ip: "", Clients: make([]*Client, 0)}
		hub.Streams[channel] = stream
	}

	ctx, cancel := context.WithCancel(context.Background())

	client := &Client{Stream: stream, Conn: conn, Send: make(chan Message), Message: make(chan Message), Channel: channel, Type: clientType}
	client.Stream.Hub.Register <- client

	go client.readPump(cancel)
	go client.writePump(ctx)

	return client, ctx, nil
}

func (c *Client) writePump(ctx context.Context) {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.Stream.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					fmt.Println(err)
				}
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				fmt.Println(err)
				return
			}

			w.Write(message.Data)

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func (c *Client) readPump(cancel context.CancelFunc) {
	defer func() {
		c.Stream.Hub.Unregister <- c
		c.Conn.Close()
		cancel()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(poingWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(poingWait)); return nil })

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println(err)
			}
			break
		}

		message := Message{c, msg, c.Channel}
		c.Message <- message
	}
}
