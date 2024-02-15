package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"
	"server/infrastructure/logger"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var _ queryservice.IMailMagazineQueryService = &MailMagazineQueryService{}

type MailMagazineQueryService struct {
	db *sql.DB
}

func NewMailMagazineQueryService() *MailMagazineQueryService {
	db := InitDB()

	return &MailMagazineQueryService{
		db: db,
	}
}

func (pq *MailMagazineQueryService) GetByID(id uuid.UUID) (*entity.MailMagazine, error) {
	mgz, err := models.FindMailMagazine(context.Background(), pq.db, id.String())
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
	return MailMagazineModelToEntity(mgz), nil
}

func (pq *MailMagazineQueryService) GetAll(pager *types.PageQuery) ([]*entity.MailMagazine, *types.PageResponse, error) {
	mgzs, err := models.MailMagazines(qm.Limit(pager.Limit()), qm.Offset(pager.Offset()), qm.OrderBy(models.MailMagazineColumns.CreateAt+" DESC")).All(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, nil
		}
		return nil, nil, err
	}
	var result []*entity.MailMagazine
	for _, mgz := range mgzs {
		result = append(result, MailMagazineModelToEntity(mgz))
	}
	count, err := models.MailMagazines().Count(context.Background(), pq.db)
	if err != nil {
		return nil, nil, err
	}

	var pageResponse *types.PageResponse = nil
	if pager != nil {
		pageResponse = types.NewPageResponse(pager, int(count))
	}
	return result, pageResponse, err
}

func MailMagazineModelToEntity(mgz *models.MailMagazine) *entity.MailMagazine {
	var prefs []entity.Prefecture
	for _, pref := range mgz.TargetPrefectures {
		prefs = append(prefs, entity.Prefecture(pref))
	}

	return &entity.MailMagazine{
		ID:                 uuid.MustParse(mgz.ID),
		Title:              mgz.Title,
		Content:            mgz.Content,
		AuthorID:           uuid.MustParse(mgz.AuthorID),
		TargetPrefecture:   &prefs,
		MailMagazineStatus: entity.MailMagazineStatus(mgz.MailMagazineStatus),
		UnsentCount:        mgz.UnsentCount,
		SentCount:          mgz.SentCount,
		SentAt:             mgz.SentAt.Ptr(),
		CreateAt:           mgz.CreateAt,
		UpdateAt:           mgz.UpdateAt,
	}
}
