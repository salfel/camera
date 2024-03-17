package handlers

import (
	"camera-server/services/database"
	"camera-server/templates"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Home(c *gin.Context) {
	db := database.GetDB()

	var session database.Session
	cookie, err := c.Cookie("session")
	if err == nil {
		db.Where("id = ?", cookie).First(&session)
	}

	var user database.User
	err = db.Model(&database.User{}).Preload("Visits").Take(&user).Error

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			templ.Handler(templates.Error(err.Error())).ServeHTTP(c.Writer, c.Request)
			return
		}
		user.Visits = nil
	}

	templ.Handler(templates.Home(user.Visits)).ServeHTTP(c.Writer, c.Request)
}
