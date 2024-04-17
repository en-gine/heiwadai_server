package entity

import (
	"time"

	"server/core/errors"

	"github.com/google/uuid"
)

const MaxStampCount = 5

type Stamp struct {
	CheckinID       *uuid.UUID
	StoreName       *string
	StoreID         *uuid.UUID
	StoreStampImage *string
	CheckInAt       *time.Time
}

type StampCard struct {
	Stamps []Stamp
}

func NewStampCard(
	userCheckIns []*Checkin,
) (*StampCard, *errors.DomainError) {
	var stamps []Stamp

	if len(userCheckIns) > MaxStampCount {
		return nil, errors.NewDomainError(errors.InvalidParameter, "チェックイン数が上限を超えています。")
	}

	// チェックインを格納
	for i := 0; i < len(userCheckIns); i++ {
		stamps[i] = Stamp{
			CheckinID:       &userCheckIns[i].ID,
			StoreName:       &userCheckIns[i].Store.Name,
			StoreID:         &userCheckIns[i].Store.ID,
			StoreStampImage: &userCheckIns[i].Store.StampImageURL,
			CheckInAt:       &userCheckIns[i].CheckInAt,
		}
	}
	return &StampCard{
		Stamps: stamps,
	}, nil
}
