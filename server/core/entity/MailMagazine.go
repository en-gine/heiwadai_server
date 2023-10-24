package entity

import (
	"time"

	"github.com/google/uuid"
)

type MailMagazine struct {
	ID                 uuid.UUID
	Title              string
	Content            string
	AuthorID           uuid.UUID
	MailMagazineStatus MailMagazineStatus
	SentCount          *int
	SentAt             *time.Time
	CreateAt           time.Time
	UpdateAt           time.Time
}

type MailMagazineStatus int

const (
	MailMagazineDraft MailMagazineStatus = iota
	MailMagazineSaved
	MailMagazineSent
)

func (s MailMagazineStatus) String() string {
	switch s {
	case MailMagazineDraft:
		return "Draft"
	case MailMagazineSaved:
		return "Saved"
	case MailMagazineSent:
		return "Sent"
	default:
		return "Unknown"
	}
}

func CreateDraftMailMagazine(
	Title string,
	Content string,
	AuthorID uuid.UUID,
) *MailMagazine {
	return &MailMagazine{
		ID:                 uuid.New(),
		Title:              Title,
		Content:            Content,
		AuthorID:           AuthorID,
		MailMagazineStatus: MailMagazineDraft,
		CreateAt:           time.Now(),
	}
}

func UpdateMailMagazine(
	ID uuid.UUID,
	Title string,
	Content string,
	AuthorID uuid.UUID,
) *MailMagazine {
	return &MailMagazine{
		ID:                 ID,
		Title:              Title,
		Content:            Content,
		AuthorID:           AuthorID,
		MailMagazineStatus: MailMagazineSaved,
		UpdateAt:           time.Now(),
	}
}

func SentMailMagazine(
	ID uuid.UUID,
	Title string,
	Content string,
	AuthorID uuid.UUID,
) *MailMagazine {
	sentAt := time.Now()
	return &MailMagazine{
		ID:                 ID,
		Title:              Title,
		Content:            Content,
		AuthorID:           AuthorID,
		MailMagazineStatus: MailMagazineSent,
		SentAt:             &sentAt,
	}
}

func RegenMailMagazine(
	ID uuid.UUID,
	Title string,
	Content string,
	AuthorID uuid.UUID,
	MailMagazineStatus MailMagazineStatus,
	SentAt *time.Time,
	CreateAt time.Time,
	UpdateAt time.Time,
) *MailMagazine {
	return &MailMagazine{
		ID:                 uuid.New(),
		Title:              Title,
		Content:            Content,
		AuthorID:           AuthorID,
		MailMagazineStatus: MailMagazineStatus,
		SentAt:             SentAt,
		CreateAt:           CreateAt,
		UpdateAt:           UpdateAt,
	}
}
