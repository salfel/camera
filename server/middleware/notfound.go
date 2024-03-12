package middleware

import (
	"camera-server/templates"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
    c.Next()

    if c.Writer.Status() == http.StatusNotFound {
        templ.Handler(templates.Error("404 Not Found")).ServeHTTP(c.Writer, c.Request)
    }
}
