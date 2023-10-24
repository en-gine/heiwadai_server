package repository

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IMessageRepository interface {
	Save(updateMessage *entity.Message) error
	Delete(messageID uuid.UUID) error
}
