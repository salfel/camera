package handlers

import (
	"camera-server/services/database"
	"camera-server/templates"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	db := database.GetDB()

	var session database.Session
	cookie, err := c.Cookie("session")
	if err == nil {
		db.Where("id = ?", cookie).First(&session)
	}

	templ.Handler(templates.Home()).ServeHTTP(c.Writer, c.Request)
}
