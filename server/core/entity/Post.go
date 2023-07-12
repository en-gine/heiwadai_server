package entity

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID         uuid.UUID
	Title      string
	Content    string
	Author     Admin
	PostStatus PostStatus
	PostDate   time.Time
	CreateAt   time.Time
	UpdateAt   time.Time
}

type PostStatus int

const (
	PostDraft PostStatus = iota
	PostPublish
)

func (s PostStatus) String() string {
	switch s {
	case PostDraft:
		return "Draft"
	case PostPublish:
		return "Publish"
	default:
		return "Unknown"
	}
}

func CreatePost(
	Title string,
	Content string,
	PostDate time.Time,
	Author Admin,
) *Post {
	return &Post{
		ID:         uuid.New(),
		Title:      Title,
		Content:    Content,
		PostStatus: PostDraft,
		PostDate:   PostDate,
		CreateAt:   time.Now(),
		Author:     Author,
	}
}

func UpdatePost(
	ID uuid.UUID,
	Title string,
	Content string,
	PostDate time.Time,
	Author Admin,
) *Post {
	return &Post{
		ID:         ID,
		Title:      Title,
		Content:    Content,
		PostStatus: PostDraft,
		PostDate:   PostDate,
		UpdateAt:   time.Now(),
		Author:     Author,
	}
}

func RegenPost(
	ID uuid.UUID,
	Title string,
	Content string,
	Author Admin,
	PostStatus PostStatus,
	PostDate time.Time,
	CreateAt time.Time,
	UpdateAt time.Time,
) *Post {
	return &Post{
		ID:         uuid.New(),
		Title:      Title,
		Content:    Content,
		Author:     Author,
		PostStatus: PostStatus,
		PostDate:   PostDate,
		UpdateAt:   UpdateAt,
		CreateAt:   CreateAt,
	}
}
