package handlers

import (
	"net/http"

	"camera-server/middleware"
	"camera-server/services/broadcast"
    "camera-server/handlers/auth"

	"github.com/gin-gonic/gin"
)

func HandleRoutes(hub *broadcast.Hub) http.Handler {
    r := gin.Default()

    r.Use(middleware.User)

    r.GET("/", Home)
    
    r.GET("video/:channel", middleware.Auth, func(c *gin.Context) {
        Video(c, hub)
    })

    r.GET("/stream/:channel", func(c *gin.Context) {
        Stream(c, hub)
    })

    a := r.Group("/auth")
    {
        a.GET("/login", auth.Login)
        a.POST("/authenticate", auth.Authenticate)

        a.GET("/register", auth.Register)
        a.POST("/create", auth.Create)

        a.POST("/logout", auth.Logout)
    }

    r.StaticFile("/js/htmx.min.js", "./public/htmx.min.js")
    r.StaticFile("/styles.css", "./public/styles.css")

    return r
}
