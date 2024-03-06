package handlers

import (
	"camera-server/components"
    . "camera-server/services"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
    db := GetDB()

    var session Session
    cookie, err := c.Cookie("session")
    if err == nil {
        db.Where("id = ?", cookie).First(&session)
    }

    templ.Handler(components.Home()).ServeHTTP(c.Writer, c.Request)
}

