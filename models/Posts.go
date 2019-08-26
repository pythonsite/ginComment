package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	User User
	Title string
	Content string
	Comments []*Comment `gorm:"-"`
}