package booking

import (
	"errors"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/infrastructure/booking/book"
	"server/infrastructure/env"
	"server/infrastructure/logger"

	uuid "github.com/google/uuid"
)

type PlanRepository struct {
	storeQuery queryservice.IStoreQueryService
}

var BookURL = env.GetEnv(env.TlbookingBookingApiUrl)

func NewPlanRepository(storeQuery queryservice.IStoreQueryService) *PlanRepository {
	return &PlanRepository{
		storeQuery: storeQuery,
	}
}

func (p *PlanRepository) GetMyBooking(userID uuid.UUID) (*[]entity.Plan, error) {
	return nil, nil
}

func (p *PlanRepository) Book(
	bookData *entity.Booking,
) error {
	reqBody := book.NewBookingRQ(
		"testID", "testPass", bookData,
	)

	res, err := Request[book.EnvelopeRQ, book.EnvelopeRS](BookURL, reqBody)
	if err != nil {
		return err
	}

	if res.Body.EntryBookingResponse.EntryBookingResult.CommonResponse.ResultCode == "False" {
		msg := res.Body.EntryBookingResponse.EntryBookingResult.CommonResponse.ErrorInfos.ErrorMsg
		code := res.Body.EntryBookingResponse.EntryBookingResult.CommonResponse.ErrorInfos.ErrorCode
		logger.Error(code + ":" + msg)
		return errors.New(msg)
	}
	return nil
}
