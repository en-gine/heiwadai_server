package entity

import "server/core/errors"

var MaxStampCount = 5

type StampCard struct {
	Checkins []*Checkin
}

func NewStampCard(
	userCheckIns []*Checkin,
) (*StampCard, *errors.DomainError) {
	stampCard := []*Checkin{}
	for i := 0; i < MaxStampCount; i++ {
		stampCard = append(stampCard, &Checkin{})
	}
	if len(userCheckIns) > MaxStampCount {
		return nil, errors.NewDomainError(errors.InvalidParameter, "チェックイン数が上限を超えています。")
	}

	copy(stampCard, userCheckIns)
	return &StampCard{
		Checkins: stampCard,
	}, nil
}
