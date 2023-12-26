package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	repository "server/core/infra/repository"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ repository.IMailMagazineRepository = &MailMagazineRepository{}

type MailMagazineRepository struct {
	db *sql.DB
}

func NewMailMagazineRepository() *MailMagazineRepository {
	db := InitDB()

	return &MailMagazineRepository{
		db: db,
	}
}

func (pq *MailMagazineRepository) Save(mailMagazine *entity.MailMagazine) error {
	var prefs []int64

	if mailMagazine.TargetPrefecture == nil {
		prefs = []int64{}
	} else {
		prefs = make([]int64, len(*mailMagazine.TargetPrefecture))
		for i, pref := range *mailMagazine.TargetPrefecture {
			prefs[i] = int64(pref)
		}
	}

	mgz := models.MailMagazine{
		ID:                 mailMagazine.ID.String(),
		Title:              mailMagazine.Title,
		Content:            mailMagazine.Content,
		AuthorID:           mailMagazine.AuthorID.String(),
		MailMagazineStatus: int(mailMagazine.MailMagazineStatus),
		CreateAt:           mailMagazine.CreateAt,
		UpdateAt:           mailMagazine.UpdateAt,
		SentAt:             null.TimeFromPtr(mailMagazine.SentAt),
		UnsentCount:        mailMagazine.UnsentCount,
		SentCount:          mailMagazine.SentCount,
		TargetPrefectures:  prefs,
	}

	err := mgz.Upsert(context.Background(), pq.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (pq *MailMagazineRepository) Delete(magazineID uuid.UUID) error {
	mgz := models.MailMagazine{
		ID: magazineID.String(),
	}
	_, err := mgz.Delete(context.Background(), pq.db)
	if err != nil {
		return err
	}
	return nil
}
