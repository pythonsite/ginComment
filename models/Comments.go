package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	UserID uint
	Content string
	PostID uint
	NickName string `gorm:"-"`
	AvatarUrl string `gorm:"-"`
	GithubUrl string `gorm:"-"`
}
