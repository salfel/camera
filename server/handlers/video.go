package handlers

import (
    "net/http"
    "fmt"

	"camera-server/services/broadcast"
    "camera-server/templates"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func Video(hub *broadcast.Hub) gin.HandlerFunc {
    return func(c *gin.Context) {
        channel := c.Param("channel")

        stream, ok := hub.Streams[channel]
        if !ok || stream.Ip == "" {
            c.Status(404)
            return
        }

        templ.Handler(templates.Video(stream.Ip)).ServeHTTP(c.Writer, c.Request)
    }
}

func Stream(hub *broadcast.Hub) gin.HandlerFunc {
    return func(c *gin.Context) {
        channel := c.Param("channel")

        client, ctx, err := broadcast.ServeWs(hub, c, channel, "camera")

        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        fmt.Println("new stream", channel)

        go broadcast.HandleVideo(client, ctx)
    }
}
