package model

import (
	"gorm.io/gorm"
	"task05/dao"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title" binding:"required"`
	Content string `gorm:"not null" json:"content" binding:"required"`
	Comment []Comment
	UserId  uint
}

func CreatePost(post *Post) (err error) {
	// 创建文章
	err = dao.DB.Create(post).Error
	if err != nil {
		return err
	}
	return nil
}

func ListPosts() ([]Post, error) {
	var posts []Post
	// 查询所有文章
	err := dao.DB.Preload("Comment").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPostById(postId uint) (Post, error) {
	var post Post
	err := dao.DB.Preload("Comment").First(&post, postId).Error
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

func DeletePost(postId uint) error {
	// 删除文章
	err := dao.DB.Delete(&Post{}, postId).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdatePost(post *Post) error {
	err := dao.DB.Model(&Post{}).Where("id = ?", post.ID).Updates(Post{
		Title:   post.Title,
		Content: post.Content,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
