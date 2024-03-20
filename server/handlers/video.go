package handlers

import (
	"net/http"

	"camera-server/services"
	"camera-server/services/broadcast"
	"camera-server/services/database"
	"camera-server/templates"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func Video(hub *broadcast.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		channel := c.Param("channel")

		// stream, ok := hub.Streams[channel]
		// if !ok || stream.Ip == "" {
		//     c.Status(404)
		//     return
		// }

		db := database.GetDB()
		ctx := c.Request.Context()
		user := ctx.Value(services.UserContext).(*database.User)

		db.Delete(&database.Visit{}, "user_id = ? AND channel = ?", user.ID, channel)

		visit := database.Visit{UserID: user.ID, Channel: channel}
		db.Create(&visit)

		templ.Handler(templates.Video("192.168.299.193")).ServeHTTP(c.Writer, c.Request)
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

		go client.HandleVideo(ctx)
	}
}
