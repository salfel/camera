package handlers

import (
	"camera-server/services"
	"camera-server/services/database"
	"camera-server/templates"
	"fmt"

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

	ctx := c.Request.Context()
	user := ctx.Value(services.UserContext).(*database.User)
	var streams []database.Stream

	err = db.Model(&user).Order("streams.updated_at desc").Association("Streams").Find(&streams)
	if err != nil {
		fmt.Println(err)
		c.Status(500)
		return
	}

	templ.Handler(templates.Home(streams)).ServeHTTP(c.Writer, c.Request)
}
