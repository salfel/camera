package handlers

import (
	"net/http"

	"camera-server/services/broadcast"

	"github.com/gin-gonic/gin"
)

func Stepper(hub *broadcast.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		channel := c.Param("channel")

		client, ctx, err := broadcast.ServeWs(hub, c, channel, "client")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		go client.HandlerStepper(ctx)
	}
}
