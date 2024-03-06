package handlers

import (
	"net/http"

	"camera-server/middleware"
	"camera-server/services/broadcast"

	"github.com/gin-gonic/gin"
)

func HandleRoutes(hub *broadcast.Hub) http.Handler {
    r := gin.Default()

    r.Use(middleware.UserMiddleware)

    r.GET("/", Home)
    
    r.GET("video/:channel", func(c *gin.Context) {
        Video(c, hub)
    }, middleware.Auth)

    r.GET("/stream/:channel", func(c *gin.Context) {
        Stream(c, hub)
    })

    r.StaticFile("/js/htmx.min.js", "./public/htmx.min.js")
    r.StaticFile("/styles.css", "./public/styles.css")

    return r
}
