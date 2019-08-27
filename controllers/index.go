package controllers

import (
	"fmt"
	"ginComment/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexGet(c *gin.Context) {

	user, _ := c.Get(models.CONTEXT_USER_KEY)
	commentlength := models.GetPostCommentsCount()
	commentuser := models.GetPostCommentUserCount()

	//models.GetCommentListMap()
	commentlistmap, _ := models.GetCommentListMap()
	fmt.Println(commentlistmap)
	for k,v := range commentlistmap {
		fmt.Printf("%#v-----%#v\n",k,v)
	}
	fmt.Println(123)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"user":user,
		"commentlistmap":commentlistmap,
		"commentlength": commentlength,
		"commentuser": commentuser,
	})
}