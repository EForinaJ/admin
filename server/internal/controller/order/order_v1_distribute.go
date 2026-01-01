package order

import (
	"context"

	v1 "server/api/order/v1"
	"server/internal/service"
)

func (c *ControllerV1) Distribute(ctx context.Context, req *v1.DistributeReq) (res *v1.DistributeRes, err error) {
	err = service.Order().CheckDistribute(ctx, req.Distribute)
	if err != nil {
		return nil, err
	}
	err = service.Order().Distribute(ctx, req.Distribute)
	return
}
