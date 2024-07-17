package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateSavedMailMagazine(t *testing.T) {
	title := "Test Mail Magazine"
	content := "This is a test mail magazine content"
	targetPrefecture := &[]Prefecture{5, 10}
	maySentCount := 1000
	authorID := uuid.New()

	mailMagazine := CreateSavedMailMagazine(title, content, targetPrefecture, maySentCount, authorID)

	assert.NotNil(t, mailMagazine)
	assert.NotEqual(t, uuid.Nil, mailMagazine.ID)
	assert.Equal(t, title, mailMagazine.Title)
	assert.Equal(t, content, mailMagazine.Content)
	assert.Equal(t, targetPrefecture, mailMagazine.TargetPrefecture)
	assert.Equal(t, maySentCount, mailMagazine.UnsentCount)
	assert.Equal(t, 0, mailMagazine.SentCount)
	assert.Equal(t, authorID, mailMagazine.AuthorID)
	assert.Equal(t, MailMagazineSaved, mailMagazine.MailMagazineStatus)
	assert.WithinDuration(t, time.Now(), mailMagazine.CreateAt, time.Second)
}

func TestUpdateMailMagazine(t *testing.T) {
	id := uuid.New()
	title := "Updated Mail Magazine"
	content := "This is an updated mail magazine content"
	targetPrefecture := &[]Prefecture{20}
	maySentCount := 500
	authorID := uuid.New()

	mailMagazine := UpdateMailMagazine(id, title, content, targetPrefecture, maySentCount, authorID)

	assert.NotNil(t, mailMagazine)
	assert.Equal(t, id, mailMagazine.ID)
	assert.Equal(t, title, mailMagazine.Title)
	assert.Equal(t, content, mailMagazine.Content)
	assert.Equal(t, targetPrefecture, mailMagazine.TargetPrefecture)
	assert.Equal(t, maySentCount, mailMagazine.UnsentCount)
	assert.Equal(t, 0, mailMagazine.SentCount)
	assert.Equal(t, authorID, mailMagazine.AuthorID)
	assert.Equal(t, MailMagazineSaved, mailMagazine.MailMagazineStatus)
	assert.WithinDuration(t, time.Now(), mailMagazine.UpdateAt, time.Second)
}

func TestCreateUnCompleteMailMagazine(t *testing.T) {
	id := uuid.New()
	title := "Uncomplete Mail Magazine"
	targetPrefecture := &[]Prefecture{20}
	unSentCount := 200
	sentCount := 300
	content := "This is an uncomplete mail magazine content"
	authorID := uuid.New()

	mailMagazine := CreateUnCompleteMailMagazine(id, title, targetPrefecture, unSentCount, sentCount, content, authorID)

	assert.NotNil(t, mailMagazine)
	assert.Equal(t, id, mailMagazine.ID)
	assert.Equal(t, title, mailMagazine.Title)
	assert.Equal(t, targetPrefecture, mailMagazine.TargetPrefecture)
	assert.Equal(t, unSentCount, mailMagazine.UnsentCount)
	assert.Equal(t, sentCount, mailMagazine.SentCount)
	assert.Equal(t, content, mailMagazine.Content)
	assert.Equal(t, authorID, mailMagazine.AuthorID)
	assert.Equal(t, MailMagazineSentUnCompleted, mailMagazine.MailMagazineStatus)
	assert.NotNil(t, mailMagazine.SentAt)
	assert.WithinDuration(t, time.Now(), *mailMagazine.SentAt, time.Second)
}

func TestCreateSentCompleteMailMagazine(t *testing.T) {
	id := uuid.New()
	title := "Complete Mail Magazine"
	targetPrefecture := &[]Prefecture{10}
	sentCount := 1000
	content := "This is a complete mail magazine content"
	authorID := uuid.New()

	mailMagazine := CreateSentCompleteMailMagazine(id, title, targetPrefecture, sentCount, content, authorID)

	assert.NotNil(t, mailMagazine)
	assert.Equal(t, id, mailMagazine.ID)
	assert.Equal(t, title, mailMagazine.Title)
	assert.Equal(t, targetPrefecture, mailMagazine.TargetPrefecture)
	assert.Equal(t, 0, mailMagazine.UnsentCount)
	assert.Equal(t, sentCount, mailMagazine.SentCount)
	assert.Equal(t, content, mailMagazine.Content)
	assert.Equal(t, authorID, mailMagazine.AuthorID)
	assert.Equal(t, MailMagazineSentCompleted, mailMagazine.MailMagazineStatus)
	assert.NotNil(t, mailMagazine.SentAt)
	assert.WithinDuration(t, time.Now(), *mailMagazine.SentAt, time.Second)
}

func TestRegenMailMagazine(t *testing.T) {
	id := uuid.New()
	title := "Regenerated Mail Magazine"
	content := "This is a regenerated mail magazine content"
	unsentCount := 100
	sentCount := 900
	targetPrefecture := &[]Prefecture{5, 10}
	authorID := uuid.New()
	status := MailMagazineSentUnCompleted
	sentAt := time.Now().Add(-1 * time.Hour)
	createAt := time.Now().Add(-2 * time.Hour)
	updateAt := time.Now().Add(-30 * time.Minute)

	mailMagazine := RegenMailMagazine(id, title, content, unsentCount, sentCount, targetPrefecture, authorID, status, &sentAt, createAt, updateAt)

	assert.NotNil(t, mailMagazine)
	assert.NotEqual(t, id, mailMagazine.ID) // RegenMailMagazine creates a new ID
	assert.Equal(t, title, mailMagazine.Title)
	assert.Equal(t, content, mailMagazine.Content)
	assert.Equal(t, unsentCount, mailMagazine.UnsentCount)
	assert.Equal(t, sentCount, mailMagazine.SentCount)
	assert.Equal(t, targetPrefecture, mailMagazine.TargetPrefecture)
	assert.Equal(t, authorID, mailMagazine.AuthorID)
	assert.Equal(t, status, mailMagazine.MailMagazineStatus)
	assert.Equal(t, &sentAt, mailMagazine.SentAt)
	assert.Equal(t, createAt, mailMagazine.CreateAt)
	assert.Equal(t, updateAt, mailMagazine.UpdateAt)
}

func TestMailMagazineStatusString(t *testing.T) {
	testCases := []struct {
		status   MailMagazineStatus
		expected string
	}{
		{MailMagazineDraft, "Draft"},
		{MailMagazineSaved, "Saved"},
		{MailMagazineSentCompleted, "SentCompleted"},
		{MailMagazineSentUnCompleted, "SentUnCompleted"},
		{MailMagazineStatus(99), "Unknown"},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, tc.status.String())
	}
}
