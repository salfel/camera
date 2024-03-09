package auth

import (
	"errors"
	"fmt"
	"time"

	"camera-server/services/database"
	"camera-server/templates"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
    templ.Handler(templates.Login()).ServeHTTP(c.Writer, c.Request)
}

func Authenticate(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    db := database.GetDB()

    var user database.User
    result := db.Where(database.User{Username: username, Password: password}).First(&user)
    if errors.Is(result.Error, gorm.ErrRecordNotFound)  {
        values := map[string]string{"username": username, "password": password }

        templ.Handler(templates.LoginForm(values, map[string]string{"password": "Wrong username or password"})).ServeHTTP(c.Writer, c.Request)
        return
    }

    session := &database.Session{UserID: user.ID}
    db.Create(&session)

    c.SetCookie("session", fmt.Sprint(session.ID), int(time.Hour * 24 * 30), "/", "localhost", false, true)

    c.Header("HX-Redirect", "/")
}
