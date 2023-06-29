package entity

import (
	"github.com/google/uuid"
)

type Post struct {
	ID         uuid.UUID
	Title      string
	Content    string
	Author     Admin
	PostStatus PostStatus
}

type PostStatus int

const (
	Draft PostStatus = iota
	Publish
)

func (s PostStatus) String() string {
	switch s {
	case Draft:
		return "Draft"
	case Publish:
		return "Publish"
	default:
		return "Unknown"
	}
}

func CreatePost(
	Title string,
	Content string,
	Author Admin,
) *Post {
	return &Post{
		ID:         uuid.New(),
		Title:      Title,
		Content:    Content,
		PostStatus: Draft,
	}
}
