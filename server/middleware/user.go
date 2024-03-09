package middleware

import (
	"camera-server/services/database"
    "context"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) *database.User {
    db := database.GetDB()

    var session database.Session

    cookie, err := c.Cookie("session")
    if err != nil {
        return nil
    }

    db.Where("id = ?", cookie).First(&session)

    if session == (database.Session{}) {
        return nil
    }

    if session.User == (database.User{}) {
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



