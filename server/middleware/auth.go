package middleware

import (
	"camera-server/services/database"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth(c *gin.Context) {
    db := database.GetDB()

    var session *database.Session
    cookie, err := c.Cookie("session")
    if err != nil {
        c.Redirect(302, "/")
        c.Abort()
        return
    } 

    err = db.Where("id = ?", cookie).First(&session).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        c.Redirect(302, "/")
        c.Abort()
    }
    
    c.Next()
}
