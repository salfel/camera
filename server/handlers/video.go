package handlers

import (
    "net/http"
    "fmt"

	"camera-server/services/broadcast"
    "camera-server/templates"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func Video(c *gin.Context, hub *broadcast.Hub) {
    channel := c.Param("channel")

    stream, ok := hub.Streams[channel]
    if !ok || stream.Ip == "" {
        c.JSON(http.StatusNotFound, "Page not found")
        return
    }

    templ.Handler(templates.Video(stream.Ip)).ServeHTTP(c.Writer, c.Request)
}

func Stream(c *gin.Context, hub *broadcast.Hub) {
    channel := c.Param("channel")

    client, ctx, err := broadcast.ServeWs(hub, c, channel, "camera")

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    fmt.Println("new stream", channel)

    go broadcast.HandleVideo(client, ctx)
}
