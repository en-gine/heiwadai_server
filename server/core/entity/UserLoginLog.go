package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserLoginLog struct {
	UserID    uuid.UUID
	LoginAt   time.Time
	RemoteID  string
	UserAgent string
}

func CreateUserLoginLog(
	userID uuid.UUID,
	remoteID string,
	userAgent string,
) *UserLoginLog {
	return &UserLoginLog{
		UserID:    userID,
		LoginAt:   time.Now(),
		RemoteID:  remoteID,
		UserAgent: userAgent,
	}
}

func RegenUserLoginLog(
	userID uuid.UUID,
	loginAt time.Time,
	remoteID string,
	userAgent string,
) *UserLoginLog {
	return &UserLoginLog{
		UserID:    userID,
		LoginAt:   loginAt,
		RemoteID:  remoteID,
		UserAgent: userAgent,
	}
}
