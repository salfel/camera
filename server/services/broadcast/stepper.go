package broadcast

import (
    "fmt"
	"context"
    "encoding/json"
)

func (client *Client) HandlerStepper(ctx context.Context) {
    for {
        select {
        case message := <-client.Message:
            if client.Type != "client" {
                continue
            }

            var msg MessageType

            err := json.Unmarshal(message.Data, &msg)
            if err != nil {
                fmt.Println(err)
                break
            }

            if msg.Type == "stepper:move" {
                moveStepper(client, message)
            }
            
        case <-ctx.Done():
            return
        }
    }
}

type moveMessage struct {
    Stepper string `json:"stepper"`
    Amount int `json:"amount"`
}

func moveStepper(c *Client, message Message) {
    var msg moveMessage
    err := json.Unmarshal(message.Data, &msg)
    if err != nil {
        fmt.Println(err)
        return
    }

    for _, client := range c.Stream.Clients {
        if client == c || client.Type != "camera" {
            continue
        }

        client.Send <- message
    }
}
