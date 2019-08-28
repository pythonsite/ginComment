package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email string `gorm:unique_index;default:null`
	Telephone string `gorm:"unique_index;defautl:null"`
	Password string `gorm:"default:null"`
	IsAdmin bool
	AvatarUrl string `gorm:"default:'/static/image/default/user-default-60x60.png'"`
	NickName string
	LockState bool `gorm:"default:'0'"`
}

func (user *User) Insert() error {
	return DB.Create(user).Error
}

func GetUser(id interface{})(*User, error) {
	var user User
	err := DB.First(&user,id).Error
	return &user,err
}

func GetUserByUserName(userName string)(*User, error) {
	var user User
	err := DB.First(&user, "email=?",userName).Error
	return &user,err
}

