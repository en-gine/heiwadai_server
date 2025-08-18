package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	repository "server/core/infra/repository"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/aarondl/sqlboiler/v4/boil"
)

var _ repository.IMessageRepository = &MessageRepository{}

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository() *MessageRepository {
	db := InitDB()

	return &MessageRepository{
		db: db,
	}
}

func (pq *MessageRepository) Save(mailMagazine *entity.Message) error {
	msg := models.Message{
		ID:          mailMagazine.ID.String(),
		Title:       mailMagazine.Title,
		Content:     mailMagazine.Content,
		DisplayDate: mailMagazine.DisplayDate,
		AuthorID:    mailMagazine.AuthorID.String(),
		CreateAt:    mailMagazine.CreateAt,
	}

	err := msg.Upsert(context.Background(), pq.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (pq *MessageRepository) Delete(messageID uuid.UUID) error {
	msg := models.Message{
		ID: messageID.String(),
	}
	_, err := msg.Delete(context.Background(), pq.db)
	if err != nil {
		return err
	}
	return nil
}
