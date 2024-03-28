package auth

import (
	"fmt"
	"time"

	"camera-server/services/database"
	"camera-server/templates"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	templ.Handler(templates.Register()).ServeHTTP(c.Writer, c.Request)
}

func Create(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	values := map[string]string{"username": username, "password": password}
	if len(password) < 5 {
		templ.Handler(templates.RegisterForm(values, map[string]string{"password": "Minimum characters is 5"})).ServeHTTP(c.Writer, c.Request)
		return
	}

	db := database.GetDB()
	if err := db.Where("username = ?", username).First(&database.User{}).Error; err == nil {

		templ.Handler(templates.RegisterForm(values, map[string]string{"username": "Username is already taken"})).ServeHTTP(c.Writer, c.Request)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}

	user := database.User{Username: username, Password: string(hash)}
	err = db.Create(&user).Error

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	session := &database.Session{UserID: user.ID}
	db.Create(&session)

	c.SetCookie("session", fmt.Sprint(session.ID), int(time.Hour*24*30), "/", "localhost", false, true)

	c.Header("HX-Redirect", "/")
	c.JSON(200, "success")
}
