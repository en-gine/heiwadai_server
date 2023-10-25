package entity

import (
	"github.com/google/uuid"
)

type UserReport struct {
	ID       uuid.UUID
	Title    string
	Content  string
	UserID   uuid.UUID
	UserName string
}

func CreateUserReport(
	title string,
	Content string,
	UserID uuid.UUID,
	UserName string,
) *UserReport {
	return &UserReport{
		ID:       uuid.New(),
		Title:    title,
		Content:  Content,
		UserID:   UserID,
		UserName: UserName,
	}
}

func RegenUserReport(
	ID uuid.UUID,
	title string,
	Content string,
	UserID uuid.UUID,
	UserName string,
) *UserReport {
	return &UserReport{
		ID:       ID,
		Title:    title,
		Content:  Content,
		UserID:   UserID,
		UserName: UserName,
	}
}
