package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Post []Post
}
type Post struct {
	gorm.Model
	Title   string
	Content string
	UserID  uint
	Comment []Comment
}
type Comment struct {
	gorm.Model
	Comment string
	PostID  uint
	UserID  uint
}
