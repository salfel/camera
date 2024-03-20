package handlers

import (
	"camera-server/templates/layouts"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func UserDropdown(c *gin.Context) {
	open := c.DefaultQuery("open", "false")
	templ.Handler(layouts.Dropdown(open != "false")).ServeHTTP(c.Writer, c.Request)
}
