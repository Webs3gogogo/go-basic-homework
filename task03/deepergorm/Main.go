package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDbInstance() (*gorm.DB, error) {
	username := "root"
	password := "admin"
	dbName := "BlogForGo"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, dbName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func main() {
	//handleCreate()
	query()
}

func handleCreate() {
	db, _ := GetDbInstance()
	db.AutoMigrate(
		&User{},
		&Post{},
		&Comment{},
	)

	user := User{Name: "Alice"}
	if err := db.Create(&user).Error; err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	posts := []Post{
		{Title: "First Post", Content: "This is the content of the first post", UserID: user.ID},
		{Title: "Second Post", Content: "This is the content of the second post", UserID: user.ID},
	}
	if err := db.Create(&posts).Error; err != nil {
		fmt.Println("Error creating post:", err)
		return
	}
	comment := []Comment{
		{Comment: "Great post!", PostID: posts[0].ID, UserID: user.ID},
		{Comment: "Thanks for sharing!", PostID: posts[0].ID, UserID: user.ID},
		{Comment: "Thanks for sharing!", PostID: posts[1].ID, UserID: user.ID},
	}
	if err := db.Create(&comment).Error; err != nil {
		fmt.Println("Error creating comment:", err)
		return
	}
}

func query() {
	db, _ := GetDbInstance()

	var user = User{}
	db.Debug().Preload("Post.Comment").Where("id=?", 1).Find(&user)
	fmt.Printf("User: %+v\n", user)

	var post = Post{}
	db.Debug().Model(&Post{}).
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("left join comments on comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count desc").
		Limit(1).
		Preload("Comment").
		First(&post)
	fmt.Printf("评论最多的文章: %s, 评论数: %d\n", post.Title, len(post.Comment))
}
func (post *Post) AfterSave(tx *gorm.DB) (err error) {
	// update the post number of user
	return
}

func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {

	var commentNum = int64(0)
	_ = tx.Model(&Comment{}).Where("post_id = ?", comment.PostID).Count(&commentNum)

	if commentNum == 0 {
		//update the state of post
	}

	return
}
