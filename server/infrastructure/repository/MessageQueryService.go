package repository

import (
	"context"
	"database/sql"
	"time"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"

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
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return MessageModelToEntity(mgz), nil
}

func (pq *MessageQueryService) GetMessagesAfter(lastCreateAt *time.Time) ([]*entity.Message, error) {
	var msgs []*models.Message

	// lastCreateAtがnilの場合は全件取得
	if lastCreateAt == nil {
		msgs, err := models.Messages().All(context.Background(), pq.db)
		if err != nil {
			return nil, err
		}
		var result []*entity.Message
		for _, mgz := range msgs {
			result = append(result, MessageModelToEntity(mgz))
		}
		return result, nil
	}

	msgs, err := models.Messages(models.MessageWhere.CreateAt.GT(*lastCreateAt)).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var result []*entity.Message
	for _, mgz := range msgs {
		result = append(result, MessageModelToEntity(mgz))
	}
	return result, nil
}

func (pq *MessageQueryService) GetAll(pager *types.PageQuery) ([]*entity.Message, error) {
	msgs, err := models.Messages(qm.Limit(pager.Limit()), qm.Offset(pager.Offset())).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var result []*entity.Message
	for _, mgz := range msgs {
		result = append(result, MessageModelToEntity(mgz))
	}
	return result, nil
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
