package entity

import "server/core/errors"

var MAX_STAMP_COUNT = 5

type StampCard struct {
	Checkins []*Checkin
}

func NewStampCard(
	userCheckIns []*Checkin,
) (*StampCard, *errors.DomainError) {
	stampCard := []*Checkin{}
	for i := 0; i < MAX_STAMP_COUNT; i++ {
		stampCard = append(stampCard, &Checkin{})
	}
	if len(userCheckIns) > MAX_STAMP_COUNT {
		return nil, errors.NewDomainError(errors.InvalidParameter, "チェックイン数が上限を超えています。")
	}

	for i, userCheckIn := range userCheckIns {
		stampCard[i] = userCheckIn
	}
	return &StampCard{
		Checkins: stampCard,
	}, nil
}
