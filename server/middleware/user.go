package middleware

import (
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

	db.Where("id = ?", cookie).Preload("User").First(&session)

	if session == nil {
		return nil
	}

	if session.User == (database.User{}) {
		return nil
	}

	return &session.User
}

func User(c *gin.Context) {
	user := getUser(c)

	ctx := context.WithValue(c.Request.Context(), "user", user)
	c.Request = c.Request.WithContext(ctx)

	c.Next()
}
