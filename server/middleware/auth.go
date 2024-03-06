package middleware

import (
	 "camera-server/services"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
    db := services.GetDB()

    var session services.Session
    cookie, err := c.Cookie("session")
    if err != nil {
        c.Redirect(302, "/")
        c.Abort()
    } else {
        db.Where("id = ?", cookie).First(&session)
    }

    if session == (services.Session{}) {
        c.Redirect(302, "/")
        c.Abort()
    }

    c.Next()
}
