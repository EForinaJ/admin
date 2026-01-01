package order

import (
	"context"

	v1 "server/api/order/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetDistributeList(ctx context.Context, req *v1.GetDistributeListReq) (res *v1.GetDistributeListRes, err error) {
	total, list, err := service.Order().GetDistributeList(ctx, req.DistributeQuery)
	if err != nil {
		return nil, err
	}

	res = &v1.GetDistributeListRes{
		Total: total,
		List:  list,
	}
	return
}
