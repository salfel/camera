package middleware

import (
	"camera-server/services"
	"camera-server/services/database"
	"context"

	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) *database.User {
	db := database.GetDB()

	var session *database.Session

	cookie, err := c.Cookie("session")
	if err != nil {
		return nil
	}

	err = db.Where("id = ?", cookie).Preload("User").First(&session).Error

	if session == nil || err != nil {
		return nil
	}

	return &session.User
}

func User(c *gin.Context) {
	user := getUser(c)

	ctx := context.WithValue(c.Request.Context(), services.UserContext, user)
	c.Request = c.Request.WithContext(ctx)

	c.Next()
}
