package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/db/models"
	"server/infrastructure/logger"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var _ queryservice.IBookQueryService = &BookQueryService{}

type BookQueryService struct {
	db *sql.DB
}

func NewBookQueryService() *BookQueryService {
	db := InitDB()

	return &BookQueryService{
		db: db,
	}
}

func (pq *BookQueryService) GetByID(bookID uuid.UUID) (*entity.Booking, error) {
	// book, err := models.FindUserBook(context.Background(), pq.db, bookID.String())
	book, err := models.UserBooks(
		models.UserBookWhere.ID.EQ(bookID.String()),
		qm.Load(models.UserBookRels.GuestDatum),
		qm.Load(models.UserBookRels.BookPlan),
	).One(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if book == nil {
		return nil, nil
	}
	guest := book.R.GuestDatum
	plan := book.R.BookPlan
	entity := BookModelToEntity(book, guest, plan)
	return entity, nil
}

func (pq *BookQueryService) GetMyBooking(userID uuid.UUID) ([]*entity.Booking, error) {
	books, err := models.UserBooks(
		models.UserBookWhere.BookUserID.EQ(userID.String()),
		qm.Load(models.UserBookRels.GuestDatum),
		qm.Load(models.UserBookRels.BookPlan)).All(context.Background(), pq.db)
	var entities []*entity.Booking
	for _, book := range books {
		guest := book.R.GuestDatum
		plan := book.R.BookPlan
		entity := BookModelToEntity(book, guest, plan)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if books == nil {
		return nil, nil
	}
	return entities, nil
}

func (pq *BookQueryService) GenerateTLBookingNumber() (*string, error) {
	var reqID string
	// Use a raw query

	query := pq.db.QueryRow("SELECT generate_booking_number()")
	err := query.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}

	err = query.Scan(&reqID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}

	return &reqID, nil
}

func BookModelToEntity(book *models.UserBook, guest *models.BookGuestDatum, plan *models.BookPlan) *entity.Booking {
	prefCode := guest.Prefecture.Ptr()
	var pref *entity.Prefecture
	if prefCode == nil {
		pref = nil
	} else {
		temp := entity.Prefecture(*prefCode)
		pref = &temp
	}

	guestEntity := &entity.GuestData{
		FirstName:     guest.FirstName,
		LastName:      guest.LastName,
		FirstNameKana: guest.FirstNameKana,
		LastNameKana:  guest.LastNameKana,
		CompanyName:   guest.CompanyName.Ptr(),
		ZipCode:       guest.ZipCode.Ptr(),
		Prefecture:    pref,
		City:          guest.City.Ptr(),
		Address:       guest.Address.Ptr(),
		Tel:           guest.Tel.Ptr(),
		Mail:          guest.Mail,
	}

	planEntity := &entity.Plan{
		ID:        plan.ID,
		Title:     plan.ID,
		Price:     uint(plan.Price),
		ImageURL:  plan.ImageURL,
		RoomType:  entity.RoomType(plan.RoomType),
		MealType:  entity.MealType{Morning: plan.MealTypeMorning, Dinner: plan.MealTypeDinner},
		SmokeType: entity.SmokeType(plan.SmokeType),
		OverView:  plan.Overview,
		StoreID:   uuid.MustParse(plan.StoreID),
	}

	return &entity.Booking{
		ID:              uuid.MustParse(book.ID),
		TlBookingNumber: book.TLBookingNumber,
		StayFrom:        book.StayFrom,
		StayTo:          book.StayTo,
		Adult:           uint(book.Adult),
		Child:           uint(book.Child),
		RoomCount:       uint(book.RoomCount),
		CheckInTime:     entity.CheckInTime(book.CheckInTime),
		TotalCost:       uint(book.TotalCost),
		GuestData:       guestEntity,
		BookPlan:        planEntity,
		BookUserID:      uuid.MustParse(book.BookUserID),
		Note:            book.Note.String,
	}
}
