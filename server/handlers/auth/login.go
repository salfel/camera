package auth

import (
    "fmt"

	"camera-server/templates"
	"camera-server/services/database"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
    templ.Handler(templates.Login()).ServeHTTP(c.Writer, c.Request)
}

func Authenticate(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    db := database.GetDB()

    fmt.Println(username, password)

    var user database.User
    err := db.Where(database.User{Username: username, Password: password}).Find(&user).Error
    if err != nil {
        fmt.Println("couldnt find user")
    }

    c.JSON(200, user)
}
