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

var _ repository.IBannerRepository = &BannerRepository{}

type BannerRepository struct {
	db *sql.DB
}

func NewBannerRepository() *BannerRepository {
	db := InitDB()

	return &BannerRepository{
		db: db,
	}
}

func (pr *BannerRepository) Save(updateBanner *entity.Banner) error {
	banner := models.Banner{
		ID:       updateBanner.ID.String(),
		ImageURL: updateBanner.ImageURL,
		URL:      updateBanner.Url,
		Status:   int(updateBanner.Status),
	}
	err := banner.Upsert(context.Background(), pr.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	return err
}

func (pr *BannerRepository) Delete(bannerId uuid.UUID) error {
	deleteBanner, err := models.FindBanner(context.Background(), pr.db, bannerId.String())
	if err != nil {
		return err
	}
	_, err = deleteBanner.Delete(context.Background(), pr.db)
	return err
}
