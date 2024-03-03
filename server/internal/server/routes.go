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

	r.GET("/", videoHandler)

    r.GET("/stream/:channel", func(c *gin.Context) {
        streamHandler(c, hub)
    })

	r.Static("/js", "./cmd/web/js")
	r.StaticFile("/styles.css", "./cmd/web/css/styles.css")

	return r
}

func videoHandler(c *gin.Context) {
    templ.Handler(web.Video()).ServeHTTP(c.Writer, c.Request)
}

func streamHandler(c *gin.Context, hub *broadcast.Hub) {
    channel := c.Param("channel")

    broadcast.ServeWs(hub, c, channel)
}
