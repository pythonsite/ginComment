package controllers

import (
	"ginComment/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK,"login.html",nil)
}

func LoginPost(c *gin.Context) {
	var (
		err error
		user *models.User
	)
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"message":"username or password can not null",
		})
		return
	}
	log.Println(username)
	user, err = models.GetUserByUserName(username)
	log.Printf("%v", user)
	log.Println(user.Password)
	log.Println(password)
	if err != nil || user.Password != password {
		log.Println("invalid username or password")
		c.HTML(http.StatusOK, "login.html", gin.H{
			"message": "invalid username or password",
		})
		return
	}
	if user.LockState {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"message":"user have bee locked",
		})
		return
	}
	s := sessions.Default(c)
	s.Clear()
	s.Set(models.SESSION_KEY, user.ID)
	s.Save()
	c.Redirect(http.StatusMovedPermanently, "/")
}


