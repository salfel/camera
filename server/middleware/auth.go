package middleware

import (
	 "camera-server/services/database"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
    db := database.GetDB()

    var session database.Session
    cookie, err := c.Cookie("session")
    if err != nil {
        c.Redirect(302, "/")
        c.Abort()
    } else {
        db.Where("id = ?", cookie).First(&session)
    }

    if session == (database.Session{}) {
        c.Redirect(302, "/")
        c.Abort()
    }

    c.Next()
}
