package middleware

import (
	"camera-server/services/database"
	"context"

	"github.com/gin-gonic/gin"
)

type ContextKey string

var userContext ContextKey = "user"

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

	ctx := context.WithValue(c.Request.Context(), userContext, user)
	c.Request = c.Request.WithContext(ctx)

	c.Next()
}
