package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	"server/core/infra/repository"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ repository.ICheckinRepository = &CheckinRepository{}

type CheckinRepository struct {
	db *sql.DB
}

func NewCheckinRepository() *CheckinRepository {
	db := InitDB()

	return &CheckinRepository{
		db: db,
	}
}

func (pr *CheckinRepository) Save(ctx context.Context, updateCheckin *entity.Checkin) error {
	checkin := models.Checkin{
		ID:        updateCheckin.ID.String(),
		StoreID:   updateCheckin.Store.ID.String(),
		UserID:    updateCheckin.User.ID.String(),
		CheckInAt: updateCheckin.CheckInAt,
		Archive:   updateCheckin.Archive,
	}
	err := checkin.Upsert(ctx, pr.db, true, []string{"id"}, boil.Infer(), boil.Infer())

	return err
}

func (pr *CheckinRepository) BulkArchive(ctx context.Context, userID uuid.UUID) error {
	ckins, err := models.Checkins(models.CheckinWhere.UserID.EQ(userID.String())).All(ctx, pr.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	_, err = ckins.UpdateAll(ctx, pr.db, models.M{models.CheckinColumns.Archive: true})
	if err != nil {
		return err
	}
	return nil
}
