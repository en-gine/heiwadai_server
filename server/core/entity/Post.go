package entity

import (
	"time"
)

type Post struct {
	ID         int
	Title      string
	Content    string
	Author     string
	PostStatus PostStatus
	PostDate   time.Time
}

type PostStatus int

func RegenPost(
	ID int,
	Title string,
	Content string,
	Author string,
	PostDate time.Time,
) *Post {
	return &Post{
		ID:       ID,
		Title:    Title,
		Content:  Content,
		Author:   Author,
		PostDate: PostDate,
	}
}
