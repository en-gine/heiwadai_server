package messaging

import (
	"server/core/entity"
	"server/core/infra/action"
)

var _ action.IAttachCouponAction = &AttachCoupon{}

type AttachCoupon struct {
}

func (ac *AttachCoupon) Issue(Coupon *entity.Coupon) (int, error) {
	return 0, nil
}
