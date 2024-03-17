package handlers

import (
	"net/http"

	"camera-server/handlers/auth"
	"camera-server/middleware"
	"camera-server/services/broadcast"

	"github.com/gin-gonic/gin"
)

func HandleRoutes(hub *broadcast.Hub) http.Handler {
	r := gin.Default()

	r.Use(middleware.User, middleware.NotFound)

	r.GET("/", Home)

	r.GET("video/:channel", middleware.Auth, Video(hub))

	r.GET("/stream/:channel", Stream(hub))

	r.GET("/stepper/:channel", Stepper(hub))

	a := r.Group("/auth")
	{
		a.GET("/login", auth.Login)
		a.POST("/authenticate", auth.Authenticate)

		a.GET("/register", auth.Register)
		a.POST("/create", auth.Create)

		a.POST("/logout", auth.Logout)
	}

	r.Static("/js", "./public/js")
	r.StaticFile("/styles.css", "./public/styles.css")

	return r
}
