package admin

import (
	"time"

	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
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

func (u *MessageUsecase) GetByID(id uuid.UUID) (*entity.Message, *errors.DomainError) {
	Message, err := u.MessageQuery.GetByID(id)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return Message, nil
}

func (u *MessageUsecase) GetList(pager *types.PageQuery) ([]*entity.Message, *types.PageResponse, *errors.DomainError) {
	Messages, pageRes, err := u.MessageQuery.GetAll(pager)
	if err != nil {
		return nil, nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return Messages, pageRes, nil
}

func (u *MessageUsecase) Create(title string, content string, postDate time.Time, autherID uuid.UUID) (*entity.Message, *errors.DomainError) {
	Message := entity.CreateMessage(title, content, postDate, autherID)

	err := u.MessageRepository.Save(Message)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return Message, nil
}

func (u *MessageUsecase) Update(title *string, content *string, postDate *time.Time, autherID uuid.UUID, MessageID uuid.UUID) (*entity.Message, *errors.DomainError) {
	oldMessage, err := u.MessageQuery.GetByID(MessageID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	if oldMessage == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "対象の投稿が見つかりません")
	}
	var updateTitle string
	var updateContent string
	var updatePostDate time.Time
	if title != nil {
		updateTitle = *title
	} else {
		updateTitle = oldMessage.Title
	}

	if content != nil {
		updateContent = *content
	} else {
		updateContent = oldMessage.Content
	}

	if postDate != nil {
		updatePostDate = *postDate
	} else {
		updatePostDate = oldMessage.DisplayDate
	}
	updateMessage := entity.RegenMessage(MessageID, updateTitle, updateContent, updatePostDate, autherID, oldMessage.CreateAt)

	err = u.MessageRepository.Save(updateMessage)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return updateMessage, nil
}

func (u *MessageUsecase) Delete(MessageID uuid.UUID) *errors.DomainError {
	deleteMessage, err := u.MessageQuery.GetByID(MessageID)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}

	if deleteMessage == nil {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "対象の投稿が見つかりません")
	}

	err = u.MessageRepository.Delete(MessageID)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return nil
}
