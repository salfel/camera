package auth

import (
	"camera-server/services"
	"camera-server/services/database"
	"camera-server/templates"
	"fmt"
	"net/http"

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
	channel := c.PostForm("channel")
	authToken := c.PostForm("authToken")

	ctx := c.Request.Context()
	user := ctx.Value(services.UserContext).(*database.User)

	db := database.GetDB()

	var stream database.Stream
	err := db.Where("channel = ?", channel).First(&stream).Error

	if err == gorm.ErrRecordNotFound {
		templ.Handler(templates.VideoForm(channel, map[string]string{"channel": "Stream " + channel + " does not exist"})).ServeHTTP(c.Writer, c.Request)
		return
	}

	var streams []database.Stream
	err = db.Model(&user).Where("channel = ?", channel).Association("Streams").Find(&streams)

	if err != nil {
		fmt.Println(err)
		c.Status(500)
		return
	}

	if len(streams) > 0 {
		c.Redirect(http.StatusSeeOther, "/video/"+channel)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(stream.AuthToken), []byte(authToken)); err != nil {
		templ.Handler(templates.VideoForm(channel, map[string]string{"authToken": "Incorrect auth token"})).ServeHTTP(c.Writer, c.Request)
		return
	}

	err = db.Model(&user).Association("Streams").Append(&stream)

	if err != nil {
		fmt.Println(err)
		c.Status(500)
		return
	}

	c.Header("HX-Redirect", "/video/"+channel)
}
