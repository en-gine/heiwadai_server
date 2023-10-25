package repository

import (
	"server/core/entity"
)

type IUserReportRepository interface {
	Save(report *entity.UserReport) error
}
