package middleware

import (
    "fmt"

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
        fmt.Println("aborted")
        return
    } 

    db.Where("id = ?", cookie).First(&session)

    fmt.Println(session)
    
    if session == (database.Session{}) {
        c.Redirect(302, "/")
        c.Abort()
    }

    c.Next()
}
