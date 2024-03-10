package auth

import (
	"camera-server/services/database"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
    sessionId, err := c.Cookie("session")
    if err != nil {
        return
    }

    c.SetCookie("session", "", -1, "/", "localhost", true, false)

    db := database.GetDB()

    db.Delete(&database.User{}, sessionId)

    c.Header("HX-Redirect", "/")
}
