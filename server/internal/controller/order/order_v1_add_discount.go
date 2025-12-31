package order

import (
	"context"

	v1 "server/api/order/v1"
	"server/internal/service"
)

func (c *ControllerV1) AddDiscount(ctx context.Context, req *v1.AddDiscountReq) (res *v1.AddDiscountRes, err error) {
	err = service.Order().CheckDiscount(ctx, req.AddDiscount)
	if err != nil {
		return nil, err
	}

	err = service.Order().AddDiscount(ctx, req.AddDiscount)
	return
}
