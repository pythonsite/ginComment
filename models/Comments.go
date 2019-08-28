package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type Comment struct {
	gorm.Model
	PostID uint	`gorm:"index:idx_post_id"`
	UserID uint  `gorm:"index:idx_user_id"`
	ReplyPk int `gorm:"index:idx_reply_pk"`  // 一级评论
	ReplyFk int `gorm:"index:idx_reply_fk"`  // 针对一级评论的回复
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
	log.Println(count)
	return
}

func GetPostCommentsCount()(count int64) {
	DB.Model(&Comment{}).Where("post_id = ?",1).Count(&count)
	return
}

func GetPostCommentUserCount()(count int64) {
	var comment = Comment{}
	DB.Model(&comment).Where("post_id=?",1).Group("user_id").Count(&count)
	return
}

func GetOneCommentALL(num uint)(commentlist []*Comment,err error) {
	rows, err := DB.Raw("select c.*, u.* from comments c inner join users u on c.user_id = u.id where c.post_id = ? and c.reply_pk=?",1,num).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var comment Comment
		DB.ScanRows(rows, &comment)
		commentlist = append(commentlist, &comment)
	}
	fmt.Printf("一级评论：%#v\n",commentlist)
	return
}

func GetTwoCommentALL(num uint)(commentlist []*Comment,err error) {
	rows, err := DB.Raw("select c.*, u.* from comments c inner join users u on c.user_id = u.id where c.reply_fk = ?", num).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var comment Comment
		DB.ScanRows(rows, &comment)
		commentlist = append(commentlist, &comment)
	}
	fmt.Printf("二级评论：%#v\n",commentlist)
	return
}



type CommentlistMap struct {
	Commentlist0 *Comment
	Commentlist []*Comment
}
func GetCommentListMap() (commentListMap []CommentlistMap,err error){
	var commentlist0 []*Comment
	var commentlist []*Comment
	var num uint = 0
	commentlist0, err = GetOneCommentALL(num)
	if err != nil {
		return
	}
	for _, v := range commentlist0 {
		commentlist,err = GetTwoCommentALL(v.ID)
		var item = CommentlistMap{}
		item.Commentlist0 = v
		item.Commentlist = commentlist
		commentListMap = append(commentListMap, item)
		commentlist = []*Comment{}
	}
	fmt.Println(commentListMap)
	return
}
