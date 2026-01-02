package order

import (
	"context"

	v1 "server/api/order/v1"
	"server/internal/service"
)

func (c *ControllerV1) Paid(ctx context.Context, req *v1.PaidReq) (res *v1.PaidRes, err error) {
	err = service.Order().CheckPaid(ctx, req.Paid)
	if err != nil {
		return nil, err
	}
	err = service.Order().Paid(ctx, req.Paid)
	return
}
