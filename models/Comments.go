package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	PostID int	`gorm:"index:idx_post_id"`
	ReplyPk int `gorm:"index:idx_reply_pk"`  // 一级评论
	ReplyFk int `gorm:"index:idx_reply_fk"`  // 针对一级评论的回复
	User User `gorm:"foreignkey:ID;association_foreignkey:user_id;auto_preload"`
	Content string  `gorm:"type:text"`
	NickName string `gorm:"-"`
	AvatarUrl string `gorm:"-"`
	GithubUrl string `gorm:"-"`
}

func (comment *Comment) Insert()error {
	return DB.Create(comment).Error
}

func GetPostComments() (count int) {
	DB.Model(&Comment{}).Where("post_id = ?",1).Where("reply_pk=?",0).Count(&count)
	return
}

func GetPostCommentsCount()(count int64) {
	DB.Model(&Comment{}).Where("post_id = ?",1).Count(&count)
	return
}

func GetPostCommentUserCount()(count int64) {
	var comment = Comment{}
	var u User
	DB.Model(&u).Where("id=?",)
	DB.Model(&comment).Where("post_id=?",1).Group("user.").Count(&count)
	return
}

func GetCommentListMap() {
	var commentlist0 []*Comment
	var commentlist []*Comment
	type CommentlistMap struct {
		Commentlist0 *Comment
		Commentlist []*Comment
	}
	var commentListMap []CommentlistMap
	DB.Model(&Comment{}).Where("post_id",1).Where("reply_pk",0).Order("created_at desc").Find(&commentlist0)
	for _, v:= range commentlist0 {
		DB.Model(&Comment{}).Where("reply_fk",v.ID).Order("created_at").Find(&commentlist)
		var item = CommentlistMap{}
		item.Commentlist0 = v
		item.Commentlist = commentlist
		commentListMap = append(commentListMap, item)
		commentlist = []*Comment{}
	}
	var commentlength int64
	var commentuser int64
	commentlength = GetPostCommentsCount()
	commentuser = GetPostCommentUserCount()
	fmt.Println("---------------")
	fmt.Println(commentListMap)
	fmt.Println(commentlength)
	fmt.Println(commentuser)
}
