package auth

import (
	"camera-server/services"
	"camera-server/services/database"
	"camera-server/templates"
	"fmt"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Video(c *gin.Context) {
	channel := c.Query("channel")

	templ.Handler(templates.VideoAuth(channel)).ServeHTTP(c.Writer, c.Request)
}

func Stream(c *gin.Context) {
	channel := c.Query("channel")
	authToken := c.PostForm("authToken")

	db := database.GetDB()

	var stream database.Stream
	err := db.Where("channel = ?", channel).First(&stream).Error
	if err == gorm.ErrRecordNotFound {
		templ.Handler(templates.VideoForm(channel, map[string]string{"channel": "Stream " + channel + " does not exist"})).ServeHTTP(c.Writer, c.Request)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(stream.AuthToken), []byte(authToken)); err != nil {
		templ.Handler(templates.VideoForm(channel, map[string]string{"authToken": "Incorrect auth token"})).ServeHTTP(c.Writer, c.Request)
		return
	}

	ctx := c.Request.Context()
	user := ctx.Value(services.UserContext).(*database.User)

	err = db.Model(&user).Association("Streams").Append(&stream)

	if err != nil {
		fmt.Println(err)
		c.Status(500)
		return
	}

	c.Header("HX-Redirect", "/video/"+channel)
}
