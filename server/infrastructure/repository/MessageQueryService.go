package repository

import (
	"context"
	"database/sql"
	"time"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"
	"server/infrastructure/logger"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var _ queryservice.IMessageQueryService = &MessageQueryService{}

type MessageQueryService struct {
	db *sql.DB
}

func NewMessageQueryService() *MessageQueryService {
	db := InitDB()

	return &MessageQueryService{
		db: db,
	}
}

func (pq *MessageQueryService) GetByID(id uuid.UUID) (*entity.Message, error) {
	mgz, err := models.FindMessage(context.Background(), pq.db, id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if mgz == nil {
		return nil, nil
	}
	return MessageModelToEntity(mgz), nil
}

func (pq *MessageQueryService) GetMessagesAfter(ID *uuid.UUID) ([]*entity.Message, error) {
	var msgs []*models.Message
	var err error
	var lastCreateAt *time.Time = nil

	if ID != nil {
		msg, err := models.FindMessage(context.Background(), pq.db, ID.String())
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			logger.Error(err.Error())
			return nil, err
		}
		if msg != nil {
			lastCreateAt = &msg.CreateAt
		}
	}
	if lastCreateAt == nil {
		msgs, err = models.Messages(qm.OrderBy(models.MessageColumns.DisplayDate+" DESC")).All(context.Background(), pq.db)
	} else {
		msgs, err = models.Messages(models.MessageWhere.CreateAt.GT(*lastCreateAt), qm.OrderBy(models.MessageColumns.DisplayDate+" DESC")).All(context.Background(), pq.db)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	var result []*entity.Message
	for _, mgz := range msgs {
		result = append(result, MessageModelToEntity(mgz))
	}
	return result, nil
}

func (pq *MessageQueryService) GetAll(pager *types.PageQuery) ([]*entity.Message, *types.PageResponse, error) {
	msgs, err := models.Messages(qm.Limit(pager.Limit()), qm.Offset(pager.Offset()), qm.OrderBy(models.MessageColumns.CreateAt+" DESC")).All(context.Background(), pq.db)
	if err != nil {
		return nil, nil, err
	}
	var result []*entity.Message
	for _, mgz := range msgs {
		result = append(result, MessageModelToEntity(mgz))
	}
	count, err := models.Messages().Count(context.Background(), pq.db)
	if err != nil {
		return nil, nil, err
	}

	var pageResponse *types.PageResponse = nil
	if pager != nil {
		pageResponse = types.NewPageResponse(pager, int(count))
	}
	return result, pageResponse, nil
}

func MessageModelToEntity(mgz *models.Message) *entity.Message {
	return &entity.Message{
		ID:          uuid.MustParse(mgz.ID),
		Title:       mgz.Title,
		Content:     mgz.Content,
		DisplayDate: mgz.DisplayDate,
		AuthorID:    uuid.MustParse(mgz.AuthorID),
		CreateAt:    mgz.CreateAt,
	}
}
