package broadcast

import (
    "context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	Stream  *Stream
	Conn    *websocket.Conn
	Send    chan Message
    Message chan Message
	Channel string
    Type    string
}

func ServeWs(hub *Hub, c *gin.Context, channel string, clientType string) (*Client, context.Context, error) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
    stream, ok := hub.Streams[channel]

    if !ok {
        stream = &Stream{Hub: hub, Ip: "", Clients: make([]*Client, 2)}
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

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

        message := Message{c, msg, c.Channel} 
        c.Message <- message
	}
}
