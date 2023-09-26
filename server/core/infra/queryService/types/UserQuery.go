package types

import "server/core/entity"

type UserQuery struct {
	FirstName     *string
	LastName      *string
	FirstNameKana *string
	LastNameKana  *string
	Prefecture    *entity.Prefecture
}
