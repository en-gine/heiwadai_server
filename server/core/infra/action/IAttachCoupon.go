package action

import (
	"server/core/entity"
)

type IAttachCouponAction interface {
	Isssue(Coupon *entity.Coupon) (int, error) // 発行数を返す
}
