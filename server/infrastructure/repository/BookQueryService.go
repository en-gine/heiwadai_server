package repository

import (
	"context"
	"database/sql"
	"time"

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
		qm.Load(models.UserBookRels.BookPlan+"."+models.BookPlanRels.PlanBookPlanStayDateInfos),
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
	dateInfos := book.R.BookPlan.R.PlanBookPlanStayDateInfos
	entity := BookModelToEntity(book, guest, plan, &dateInfos)
	return entity, nil
}

func (pq *BookQueryService) GetMyBooking(userID uuid.UUID) ([]*entity.Booking, error) {
	books, err := models.UserBooks(
		models.UserBookWhere.BookUserID.EQ(userID.String()),
		qm.Load(models.UserBookRels.GuestDatum),
		qm.Load(models.UserBookRels.BookPlan),
		qm.Load(models.UserBookRels.BookPlan+"."+models.BookPlanRels.PlanBookPlanStayDateInfos),
		models.UserBookWhere.StayFrom.GT(time.Now().AddDate(0, 0, -1)),
		models.UserBookWhere.DelateAt.IsNull(),
		qm.OrderBy(models.UserBookColumns.StayFrom+" ASC"),
	).All(context.Background(), pq.db)
	var entities []*entity.Booking
	for _, book := range books {
		guest := book.R.GuestDatum
		plan := book.R.BookPlan
		dateInfos := book.R.BookPlan.R.PlanBookPlanStayDateInfos
		entity := BookModelToEntity(book, guest, plan, &dateInfos)
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

func (pq *BookQueryService) GenerateBookDataID() (*string, error) {
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

func BookModelToEntity(book *models.UserBook, guest *models.BookGuestDatum, plan *models.BookPlan, dateInfos *models.BookPlanStayDateInfoSlice) *entity.Booking {
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

	var dateInfosEntity []entity.StayDateInfo

	for _, dateInfo := range *dateInfos {
		if dateInfo == nil {
			continue
		}
		dateInfosEntity = append(dateInfosEntity, entity.StayDateInfo{
			StayDate:           dateInfo.StayDate,
			StayDateTotalPrice: uint(dateInfo.StayDateTotalPrice),
		})
	}
	planEntity := entity.RegenPlan(
		plan.ID,
		plan.Title,
		uint(plan.Price),
		plan.ImageURL,
		entity.RoomType(plan.RoomType),
		entity.MealType{Morning: plan.MealTypeMorning, Dinner: plan.MealTypeDinner},
		entity.SmokeType(plan.SmokeType),
		plan.Overview,
		uuid.MustParse(plan.StoreID),
		plan.TLBookdataRoomTypeCode,
		plan.TLBookdataRoomTypeName,
	)

	bookPlan := &entity.PlanStayDetail{
		Plan:          planEntity,
		StayDateInfos: &dateInfosEntity,
	}

	return entity.RegenBooking(
		uuid.MustParse(book.ID),
		book.StayFrom,
		book.StayTo,
		uint(book.Adult),
		uint(book.Child),
		uint(book.RoomCount),
		entity.CheckInTime(book.CheckInTime),
		uint(book.TotalCost),
		guestEntity,
		bookPlan,
		uuid.MustParse(book.BookUserID),
		book.Note.String,
		book.TLBookdataID,
		&book.TLBookingNumber,
	)
}
