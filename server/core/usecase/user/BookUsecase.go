package user

import (
	"time"

	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
	"server/infrastructure/logger"

	"github.com/google/uuid"
)

type BookUsecase struct {
	bookQuery queryservice.IBookQueryService
	bookRepo  repository.IBookRepository
	bookAPI   repository.IBookAPIRepository
}

func NewBookUsecase(
	bookQuery queryservice.IBookQueryService,
	bookRepo repository.IBookRepository,
	bookAPI repository.IBookAPIRepository,
) *BookUsecase {
	return &BookUsecase{
		bookQuery: bookQuery,
		bookRepo:  bookRepo,
		bookAPI:   bookAPI,
	}
}

func (u *BookUsecase) GetMyBook(userID uuid.UUID) ([]*entity.Booking, *errors.DomainError) {
	books, err := u.bookQuery.GetMyBooking(userID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return books, nil
}

func (u *BookUsecase) Cancel(bookID uuid.UUID) *errors.DomainError {
	book, err := u.bookQuery.GetByID(bookID)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}

	if book == nil {
		return errors.NewDomainError(errors.QueryError, "該当の予約が存在しません。")
	}

	err = u.bookAPI.Cancel(book)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, "キャンセル処理がAPIレベルで失敗しました。")
	}

	err = u.bookRepo.Delete(bookID)
	if err != nil {
		logger.Errorf("キャンセル処理がDBレベルで失敗しました。%s", bookID.String())
		return errors.NewDomainError(errors.QueryError, "キャンセル処理がDBレベルで失敗しました。")
	}
	return nil
}

func (u *BookUsecase) GetByID(bookID uuid.UUID) (*entity.Booking, *errors.DomainError) {
	book, err := u.bookQuery.GetByID(bookID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	if book == nil {
		return nil, errors.NewDomainError(errors.QueryError, "該当の予約が存在しません。")
	}
	return book, nil
}

func (u *BookUsecase) Reserve(
	stayFrom time.Time,
	stayTo time.Time,
	adult uint,
	child uint,
	roomCount uint,
	CheckInTime entity.CheckInTime,
	TotalCost uint,
	GuestData *entity.GuestData,
	BookPlan *entity.Plan,
	BookUserID uuid.UUID,
	Note string,
) *errors.DomainError {
	newBook := entity.CreateBooking(
		stayFrom,
		stayTo,
		adult,
		child,
		roomCount,
		CheckInTime,
		TotalCost,
		GuestData,
		BookPlan,
		BookUserID,
		Note,
		"",
	)
	TLBookingNumber, err := u.bookAPI.Reserve(newBook)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	if TLBookingNumber == nil || *TLBookingNumber != "" {
		return errors.NewDomainError(errors.RepositoryError, "TLBookingNumberが存在しません。処理を中止します。")
	}

	newBook.TlBookingNumber = *TLBookingNumber

	err = u.bookRepo.Save(newBook)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}
