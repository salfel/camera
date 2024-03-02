package server

import (
	"camera-server/cmd/web"
	"camera-server/internal/server/broadcast"
	"github.com/a-h/templ"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes(hub *broadcast.Hub) http.Handler {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		templ.Handler(web.Base()).ServeHTTP(c.Writer, c.Request)
	})

	r.Static("/js", "./cmd/web/js")
	r.StaticFile("/styles.css", "./cmd/web/css/styles.css")

	return r
}
