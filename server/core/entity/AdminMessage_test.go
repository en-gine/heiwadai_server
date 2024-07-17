package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateMessage(t *testing.T) {
	title := "Test Message"
	content := "This is a test message content."
	displayDate := time.Now().Add(24 * time.Hour) // 明日の日付
	authorID := uuid.New()

	message := CreateMessage(title, content, displayDate, authorID)

	assert.NotNil(t, message)
	assert.NotEqual(t, uuid.Nil, message.ID)
	assert.Equal(t, title, message.Title)
	assert.Equal(t, content, message.Content)
	assert.Equal(t, displayDate, message.DisplayDate)
	assert.Equal(t, authorID, message.AuthorID)
	assert.WithinDuration(t, time.Now(), message.CreateAt, time.Second)
}

func TestRegenMessage(t *testing.T) {
	id := uuid.New()
	title := "Regen Message"
	content := "This is a regenerated message content."
	displayDate := time.Now().Add(48 * time.Hour) // 明後日の日付
	authorID := uuid.New()
	createAt := time.Now().Add(-24 * time.Hour) // 昨日の日付

	message := RegenMessage(id, title, content, displayDate, authorID, createAt)

	assert.NotNil(t, message)
	assert.Equal(t, id, message.ID)
	assert.Equal(t, title, message.Title)
	assert.Equal(t, content, message.Content)
	assert.Equal(t, displayDate, message.DisplayDate)
	assert.Equal(t, authorID, message.AuthorID)
	assert.Equal(t, createAt, message.CreateAt)
}
