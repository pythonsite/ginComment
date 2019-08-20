package controllers

import (
	"ginComment/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexGet(c *gin.Context) {
	user, _ := c.Get(models.CONTEXT_USER_KEY)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"user":user,
	})
}