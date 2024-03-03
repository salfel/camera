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

	r.GET("/", VideoHandler)

	r.Static("/js", "./cmd/web/js")
	r.StaticFile("/styles.css", "./cmd/web/css/styles.css")

	return r
}

func VideoHandler(c *gin.Context) {
    templ.Handler(web.Video()).ServeHTTP(c.Writer, c.Request)
}
