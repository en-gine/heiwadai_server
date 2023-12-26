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
	TargetPrefecture   *[]Prefecture
	MailMagazineStatus MailMagazineStatus
	UnsentCount        int
	SentCount          int
	SentAt             *time.Time
	CreateAt           time.Time
	UpdateAt           time.Time
}

type MailMagazineLog struct {
	MailMagazineID uuid.UUID
	UserID         uuid.UUID
	Email          string
	SentAt         *time.Time
}

type MailMagazineStatus int

const (
	MailMagazineDraft MailMagazineStatus = iota
	MailMagazineSaved
	MailMagazineSentCompleted
	MailMagazineSentUnCompleted
)

func (s MailMagazineStatus) String() string {
	switch s {
	case MailMagazineDraft:
		return "Draft"
	case MailMagazineSaved:
		return "Saved"
	case MailMagazineSentCompleted:
		return "SentCompleted"
	case MailMagazineSentUnCompleted:
		return "SentUnCompleted"
	default:
		return "Unknown"
	}
}

func CreateSavedMailMagazine(
	Title string,
	Content string,
	TargetPrefecture *[]Prefecture,
	MaySentCount int,
	AuthorID uuid.UUID,
) *MailMagazine {
	return &MailMagazine{
		ID:                 uuid.New(),
		Title:              Title,
		Content:            Content,
		UnsentCount:        MaySentCount,
		SentCount:          0,
		AuthorID:           AuthorID,
		MailMagazineStatus: MailMagazineDraft,
		CreateAt:           time.Now(),
	}
}

func UpdateMailMagazine(
	ID uuid.UUID,
	Title string,
	Content string,
	TargetPrefecture *[]Prefecture,
	MaySentCount int,
	AuthorID uuid.UUID,
) *MailMagazine {
	return &MailMagazine{
		ID:                 ID,
		Title:              Title,
		Content:            Content,
		UnsentCount:        MaySentCount,
		SentCount:          0,
		AuthorID:           AuthorID,
		MailMagazineStatus: MailMagazineSaved,
		UpdateAt:           time.Now(),
	}
}

func CreateUnCompleteMailMagazine(
	ID uuid.UUID,
	Title string,
	TargetPrefecture *[]Prefecture,
	UnSentCount int,
	SentCount int,
	Content string,
	AuthorID uuid.UUID,
) *MailMagazine {
	sentAt := time.Now()
	return &MailMagazine{
		ID:                 ID,
		Title:              Title,
		UnsentCount:        UnSentCount,
		Content:            Content,
		SentCount:          SentCount,
		AuthorID:           AuthorID,
		MailMagazineStatus: MailMagazineSentUnCompleted,
		SentAt:             &sentAt,
	}
}

func CreateSentCompleteMailMagazine(
	ID uuid.UUID,
	Title string,
	TargetPrefecture *[]Prefecture,
	SentCount int,
	Content string,
	AuthorID uuid.UUID,
) *MailMagazine {
	sentAt := time.Now()
	return &MailMagazine{
		ID:                 ID,
		Title:              Title,
		UnsentCount:        0,
		Content:            Content,
		SentCount:          SentCount,
		AuthorID:           AuthorID,
		MailMagazineStatus: MailMagazineSentCompleted,
		SentAt:             &sentAt,
	}
}

func RegenMailMagazine(
	ID uuid.UUID,
	Title string,
	Content string,
	UnsentCount int,
	SentCount int,
	TargetPrefecture *[]Prefecture,
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
		UnsentCount:        UnsentCount,
		SentCount:          SentCount,
		MailMagazineStatus: MailMagazineStatus,
		SentAt:             SentAt,
		CreateAt:           CreateAt,
		UpdateAt:           UpdateAt,
	}
}
