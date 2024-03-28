package handlers

import (
	"fmt"
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

		strmn, ok := hub.Streams[channel]
		if !ok || strmn.Ip == "" {
			c.Status(404)
			return
		}

		db := database.GetDB()
		ctx := c.Request.Context()
		user := ctx.Value(services.UserContext).(*database.User)

		var streams []database.Stream
		err := db.Model(&user).Where("channel = ?", channel).Association("Streams").Find(&streams)

		if err != nil {
			fmt.Println(err)
			c.Status(500)
			return
		}

		if len(streams) == 0 {
			c.Redirect(http.StatusSeeOther, "/video/"+channel+"/auth")
			return
		}

		templ.Handler(templates.Video()).ServeHTTP(c.Writer, c.Request)
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
