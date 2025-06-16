package model

import (
	"gorm.io/gorm"
	"log"
	"task05/dao"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content" binding:"required"`
	UserID  uint   `json:"userId"`
	PostID  uint   `json:"postId"`
}

func GetAllCommentsByPostId(postId uint) ([]Comment, error) {
	var comments []Comment
	// 查询指定文章的所有评论
	err := dao.DB.Where("post_id = ?", postId).Find(&comments).Error
	if err != nil {
		log.Fatal("Error fetching comments:", err)
		return nil, err
	}
	return comments, nil
}

func CreateComment(comment *Comment) error {
	// 创建评论
	err := dao.DB.Create(comment).Error
	if err != nil {
		return err
	}
	return nil
}
