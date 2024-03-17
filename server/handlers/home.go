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

	var visits []database.Visit
	err = db.Where("user_id = ?", session.UserID).Order("created_at desc").Find(&visits).Error

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			templ.Handler(templates.Error(err.Error())).ServeHTTP(c.Writer, c.Request)
			return
		}
	}

	templ.Handler(templates.Home(visits)).ServeHTTP(c.Writer, c.Request)
}
