package broadcast

import (
	"camera-server/services/database"
	"context"
	"encoding/json"
	"fmt"
	"net"

	rtsptowebrtc "github.com/salfel/RTSPtoWebRTC"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Stream struct {
	Hub          *Hub
	Ip           string
	AuthToken    string
	Clients      []*Client
	XOrientation int
	YOrientation int
}

type registerMessage struct {
	Ip        string `json:"ip"`
	AuthToken string `json:"authToken"`
}

func (client *Client) HandleVideo(ctx context.Context) {
	for {
		select {
		case message := <-client.Message:
			if client.Type != "camera" {
				continue
			}

			var msg MessageType
			err := json.Unmarshal(message.Data, &msg)
			if err != nil {
				fmt.Println(err)
				break
			}

			if msg.Type == "register:ip" {
				client.registerIp(message)
			}

		case <-ctx.Done():
			rtsptowebrtc.RemoveStream(client.Channel)

			delete(client.Stream.Hub.Streams, client.Channel)
			return
		}
	}
}
func (client *Client) registerIp(message Message) {
	db := database.GetDB()

	var msg registerMessage
	err := json.Unmarshal(message.Data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	if client.Stream.Ip != "" {
		return
	}

	ip := net.ParseIP(msg.Ip)
	if ip == nil {
		client.Send <- Message{Data: []byte("Invalid Ip"), Channel: client.Channel}
		return
	}

	client.Stream.Ip = msg.Ip
	client.Stream.AuthToken = msg.AuthToken
	client.Send <- Message{Data: []byte(client.Stream.Ip), Channel: client.Channel}

	rtsptowebrtc.ServeStream(client.Channel, rtsptowebrtc.StreamST{
		OnDemand:     false,
		DisableAudio: true,
		URL:          "rtsp://" + client.Stream.Ip + ":8554/cam",
	})

	var stream database.Stream
	err = db.Where("channel = ?", client.Channel).First(&stream).Error
	if err != gorm.ErrRecordNotFound {
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(client.Stream.AuthToken), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Create(&database.Stream{
		Channel:   client.Channel,
		AuthToken: string(hash),
	})
}
