package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	ObjPk *Post
	ReplyPk int64
	ReplyFk int64
	User *User
	Content string
	PostID uint
	NickName string `gorm:"-"`
	AvatarUrl string `gorm:"-"`
	GithubUrl string `gorm:"-"`
}

func (comment *Comment) Insert()error {
	return DB.Create(comment).Error
}

