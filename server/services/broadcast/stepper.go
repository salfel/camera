package broadcast

import (
	"camera-server/services/database"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"time"
)

const MIN_X = 0
const MAX_X = 100
const MIN_Y = 0
const MAX_Y = 100

func (client *Client) HandleStepper(ctx context.Context) {
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
				client.moveStepper(message)
			}

		case <-ctx.Done():
			return
		}
	}
}

type moveMessage struct {
	Axis   string `json:"axis"`
	Amount int    `json:"amount"`
}

func (c *Client) moveStepper(message Message) {
	var msg moveMessage
	err := json.Unmarshal(message.Data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.updateAmount(msg)

	for _, client := range c.Stream.Clients {
		if client == c || client.Type != "camera" {
			continue
		}

		client.Send <- message
	}
}

func (c *Client) updateAmount(msg moveMessage) {
	if msg.Axis == "x" {
		c.Stream.XOrientation = int(math.Min(MAX_X, math.Max(MIN_X, float64(c.Stream.XOrientation+msg.Amount))))
	} else if msg.Axis == "y" {
		c.Stream.YOrientation = int(math.Min(MAX_Y, math.Max(MIN_Y, float64(c.Stream.YOrientation+msg.Amount))))
	}
}

func (c *Client) ListenForOrientation() {
	ticker := time.NewTicker(time.Second)
	lastX := c.Stream.XOrientation
	lastY := c.Stream.YOrientation
	for range ticker.C {
		if lastX != c.Stream.XOrientation || lastY != c.Stream.YOrientation {
			lastX = c.Stream.XOrientation
			lastY = c.Stream.YOrientation

			c.StoreOrientationInDB()
		}
	}
}

func (c *Client) StoreOrientationInDB() {
	db := database.GetDB()
	var stream database.Stream

	db.Where("channel = ?", c.Channel).First(&stream)

	stream.XOrientation = c.Stream.XOrientation
	stream.YOrientation = c.Stream.YOrientation
	db.Save(&stream)
}
