package broadcast

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	rtsptowebrtc "github.com/salfel/RTSPtoWebRTC"
)

type Stream struct {
	Hub     *Hub
	Ip      string
	Clients []*Client
}

type registerMessage struct {
	Ip string `json:"ip"`
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

				client.Stream.Ip = msg.Ip
				client.Send <- Message{Data: []byte(client.Stream.Ip), Channel: client.Channel}

				rtsptowebrtc.ServeStream(client.Channel, rtsptowebrtc.StreamST{
					OnDemand:     false,
					DisableAudio: true,
					URL:          "rtsp://" + client.Stream.Ip + ":8554/cam",
				})
			}
		case <-ctx.Done():
			rtsptowebrtc.RemoveStream(client.Channel)

			delete(client.Stream.Hub.Streams, client.Channel)
			return
		}
	}
}
