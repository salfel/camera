package broadcast

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
)

type Stream struct {
    Hub     *Hub
    Ip      string 
    Clients []*Client
}

type messageType struct {
    Type string `json:"type"`
}

type registerMessage struct {
    Type    string `json:"type"`
    Ip      string `json:"ip"` 
}

func HandleVideo(client *Client, ctx context.Context) {
    for {
        select {
        case message := <-client.Stream.Hub.Broadcast:
            if message.Sender == client && client.Type == "camera" {
                var msg messageType
                err := json.Unmarshal(message.Data, &msg)
                if err != nil {
                    fmt.Println(err)
                    break
                }

                if msg.Type == "register:ip" {
                    var msg registerMessage
                    err := json.Unmarshal(message.Data, &msg)
                    if err != nil {
                        fmt.Println(err)
                        break
                    }

                    if client.Stream.Ip != "" {
                        break
                    }

                    ip := net.ParseIP(msg.Ip)
                    if ip == nil {
                        client.Send <- Message{Data: []byte("Invalid Ip"), Channel: client.Channel}
                        break
                    }

                    fmt.Println(msg.Ip)
                    client.Stream.Ip = msg.Ip
                    client.Send <- Message{Data: []byte(client.Stream.Ip), Channel: client.Channel}
                }
            }
        case <-ctx.Done():
            return
        }
    }
}
