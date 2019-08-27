package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	UserID uint `gorm:"index"`
	Title string
	Content string
	Comments []*Comment `gorm:"-"`
}