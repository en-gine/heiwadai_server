package entity

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID          uuid.UUID
	Title       string
	Content     string
	DisplayDate time.Time
	AuthorID    uuid.UUID
	CreateAt    time.Time
}

func CreateMessage(
	title string,
	Content string,
	DisplayDate time.Time,
	AuthorID uuid.UUID,
) *Message {
	return &Message{
		ID:          uuid.New(),
		Title:       title,
		Content:     Content,
		DisplayDate: DisplayDate,
		AuthorID:    AuthorID,
		CreateAt:    time.Now(),
	}
}

func RegenMessage(
	ID uuid.UUID,
	title string,
	Content string,
	DisplayDate time.Time,
	AuthorID uuid.UUID,
	CreateAt time.Time,
) *Message {
	return &Message{
		ID:          ID,
		Title:       title,
		Content:     Content,
		DisplayDate: DisplayDate,
		AuthorID:    AuthorID,
		CreateAt:    CreateAt,
	}
}
