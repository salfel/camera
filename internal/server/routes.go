package server

import (
	"camera-server/internal/server/broadcast"
	"net/http"

	"github.com/gin-gonic/gin"

	"camera-server/cmd/web"
	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes(hub *broadcast.Hub) http.Handler {
	r := gin.Default()

	r.GET("/", s.HelloWorldHandler)

	r.Static("/js", "./cmd/web/js")

	r.GET("/web", func(c *gin.Context) {
		templ.Handler(web.HelloForm()).ServeHTTP(c.Writer, c.Request)
	})

	r.POST("/hello", func(c *gin.Context) {
		web.HelloWebHandler(c.Writer, c.Request)
	})

	r.GET("/video", func(c *gin.Context) {
		broadcast.ServeWs(hub, c.Writer, c.Request)
	})

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}
