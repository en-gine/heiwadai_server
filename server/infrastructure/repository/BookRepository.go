package repository

import (
	"context"
	"database/sql"
	"time"

	"server/core/entity"
	repository "server/core/infra/repository"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ repository.IBookRepository = &BookRepository{}

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository() *BookRepository {
	db := InitDB()

	return &BookRepository{
		db: db,
	}
}

func (pq *BookRepository) Save(entity *entity.Booking) error {
	guest := models.BookGuestDatum{
		ID:            uuid.New().String(),
		FirstName:     entity.GuestData.FirstName,
		LastName:      entity.GuestData.LastName,
		FirstNameKana: entity.GuestData.FirstNameKana,
		LastNameKana:  entity.GuestData.LastNameKana,
		CompanyName:   null.StringFromPtr(entity.GuestData.CompanyName),
		ZipCode:       null.StringFromPtr(entity.GuestData.ZipCode),
		Prefecture:    null.IntFromPtr((*int)(entity.GuestData.Prefecture)),
		City:          null.StringFromPtr(entity.GuestData.City),
		Address:       null.StringFromPtr(entity.GuestData.Address),
		Tel:           null.StringFromPtr(entity.GuestData.Tel),
		Mail:          entity.GuestData.Mail,
	}

	plan := models.BookPlan{
		ID:                     uuid.New().String(),
		PlanID:                 entity.BookPlan.Plan.ID,
		Title:                  entity.BookPlan.Plan.Title,
		Price:                  int(entity.BookPlan.Plan.Price),
		ImageURL:               entity.BookPlan.Plan.ImageURL,
		RoomType:               int(entity.BookPlan.Plan.RoomType),
		MealTypeMorning:        entity.BookPlan.Plan.MealType.Morning,
		MealTypeDinner:         entity.BookPlan.Plan.MealType.Dinner,
		SmokeType:              int(entity.BookPlan.Plan.SmokeType),
		Overview:               entity.BookPlan.Plan.OverView,
		StoreID:                entity.BookPlan.Plan.StoreID.String(),
		TLBookdataRoomTypeCode: entity.BookPlan.Plan.TlBookingRoomTypeCode,
		TLBookdataRoomTypeName: entity.BookPlan.Plan.TlBookingRoomTypeName,
	}

	var dateInfos []models.BookPlanStayDateInfo
	for _, dateInfo := range *entity.BookPlan.StayDateInfos {
		dateInfos = append(dateInfos, models.BookPlanStayDateInfo{
			PlanID:             plan.ID,
			StayDate:           dateInfo.StayDate,
			StayDateTotalPrice: int(dateInfo.StayDateTotalPrice),
		})
	}

	book := models.UserBook{
		ID:              entity.ID.String(),
		TLBookdataID:    entity.TlDataID,
		TLBookingNumber: *entity.TlBookingNumber,
		StayFrom:        entity.StayFrom,
		StayTo:          entity.StayTo,
		Adult:           int(entity.Adult),
		Child:           int(entity.Child),
		RoomCount:       int(entity.RoomCount),
		CheckInTime:     entity.CheckInTime.String(),
		TotalCost:       int(entity.TotalCost),
		GuestDataID:     guest.ID,
		BookPlanID:      plan.ID,
		BookUserID:      entity.BookUserID.String(),
		Note:            null.StringFrom(entity.Note),
	}
	ctx := context.Background()
	tran, err := pq.db.BeginTx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tran.Rollback()
		} else {
			_ = tran.Commit()
		}
	}()

	if err != nil {
		return err
	}

	err = plan.Upsert(ctx, tran, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}

	for _, dateInfo := range dateInfos {
		err = dateInfo.Upsert(ctx, tran, true, []string{"plan_id", "stay_date"}, boil.Infer(), boil.Infer())
		if err != nil {
			return err
		}
	}

	err = guest.Upsert(ctx, tran, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}

	err = book.Upsert(ctx, tran, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (pq *BookRepository) Delete(bookID uuid.UUID) error {
	book, err := models.FindUserBook(context.Background(), pq.db, bookID.String())
	if err != nil {
		return err
	}
	_, err = book.Delete(context.Background(), pq.db)
	if err != nil {
		return err
	}
	return nil
}

func (pq *BookRepository) SoftDelete(bookID uuid.UUID) error {
	book, err := models.FindUserBook(context.Background(), pq.db, bookID.String())
	if err != nil {
		return err
	}
	book.DelateAt = null.TimeFrom(time.Now())
	_, err = book.Update(context.Background(), pq.db, boil.Whitelist(models.UserBookColumns.DelateAt))
	if err != nil {
		return err
	}
	return nil
}
