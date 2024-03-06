package middleware

import (
	"camera-server/services"
    "context"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) *services.User {
    db := services.GetDB()

    var session services.Session

    cookie, err := c.Cookie("session")
    if err != nil {
        return nil
    }

    db.Where("id = ?", cookie).First(&session)

    if session == (services.Session{}) {
        return nil
    }

    if session.User == (services.User{}) {
        return nil
    }

    return &session.User
}

func UserMiddleware(c *gin.Context) {
    user := GetUser(c)

    ctx := context.WithValue(c.Request.Context(), "user", user)
    c.Request = c.Request.WithContext(ctx)

    c.Next()
}



