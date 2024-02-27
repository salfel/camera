package server

import (
	"camera-server/internal/server/broadcast"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes(hub *broadcast.Hub) http.Handler {
	r := gin.Default()

	r.Static("/js", "./cmd/web/js")
	r.StaticFile("/styles.css", "./cmd/web/css/styles.css")

	return r
}
