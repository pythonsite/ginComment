package controllers

import (
	"ginComment/models"
	"ginComment/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func CommentAdd(c *gin.Context) {
	if c.Request.Method == "POST"{
		var comment models.Comment
		postID,_ := strconv.Atoi(strings.TrimSpace(c.GetString("post_id")))
		replyPK := strings.TrimSpace(c.GetString("reply_pk"))
		replyFK := strings.TrimSpace(c.GetString("reply_fk"))
		comment_content := strings.TrimSpace(c.GetString("comment_content"))
		security_hash := strings.TrimSpace(c.GetString("security_hash"))
		timestamp := strings.TrimSpace(c.GetString("timestamp"))
		Out := map[string]string{"err":"false","msg":"请求参数不合法"}
		if comment_content != "" && security_hash != "" {
			strcount := len([]rune(comment_content))
			if strcount < 169 {
				checkstr := utils.MD5([]byte(replyPK + timestamp))
				if checkstr == security_hash {
					comment.Content = comment_content
				}
			}
		}
	}
}