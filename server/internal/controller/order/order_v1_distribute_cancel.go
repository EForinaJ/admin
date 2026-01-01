package order

import (
	"context"

	v1 "server/api/order/v1"
	"server/internal/service"
)

func (c *ControllerV1) DistributeCancel(ctx context.Context, req *v1.DistributeCancelReq) (res *v1.DistributeCancelRes, err error) {
	err = service.Order().CheckDistributeCancel(ctx, req.DistributeCancel)
	if err != nil {
		return nil, err
	}
	err = service.Order().DistributeCancel(ctx, req.DistributeCancel)
	return
}
