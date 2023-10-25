package user

import (
	"time"

	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"

	"github.com/google/uuid"
)

type MessageUsecase struct {
	MessageRepository repository.IMessageRepository
	MessageQuery      queryservice.IMessageQueryService
}

func NewMessageUsecase(MessageRepository repository.IMessageRepository, MessageQuery queryservice.IMessageQueryService) *MessageUsecase {
	return &MessageUsecase{
		MessageRepository: MessageRepository,
		MessageQuery:      MessageQuery,
	}
}

func (u *MessageUsecase) GetAfter(ID *uuid.UUID) ([]*entity.Message, *errors.DomainError) {
	var lastDate *time.Time = nil
	if ID != nil {
		msg, err := u.MessageQuery.GetByID(*ID)
		if err != nil {
			return nil, errors.NewDomainError(errors.QueryError, err.Error())
		}
		if msg != nil {
			lastDate = &msg.CreateAt
		}
	}

	msgs, err := u.MessageQuery.GetMessagesAfter(lastDate)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return msgs, nil
}
