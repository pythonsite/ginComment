package main

import (
	"ginComment/controllers"
	"ginComment/models"
	_ "ginComment/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	gin.SetMode(models.Conf.RunMode)
	router := gin.Default()
	setSessions(router)
	router.Use(ShareData())
	router.LoadHTMLGlob(filepath.Join(getCurrentDirectory(), "./views/*"))

	router.Static("/static", filepath.Join(getCurrentDirectory(), "./static"))
	router.GET("/", controllers.IndexGet)
	router.GET("/login",controllers.LoginGet)
	router.POST("/login",controllers.LoginPost)

	err := router.Run(models.Conf.General.Addr)
	if err != nil {
		log.Fatal(err)
	}

}

func setSessions(router *gin.Engine) {
	store := cookie.NewStore([]byte(models.Conf.General.SessionSecret))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 2 * 86400, Path: "/"})
	router.Use(sessions.Sessions("gin-session",store))
}

func ShareData() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if uID := session.Get(models.SESSION_KEY);uID != nil {
			user,err := models.GetUser(uID)
			if err == nil {
				c.Set(models.CONTEXT_USER_KEY, user)
			}
			if models.Conf.General.SignupEnabled {
				c.Set("SignupEnabled", true)
			}
			c.Next()
		}
	}
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir,"\\", "/", -1)
}