package order

import (
	"context"

	v1 "server/api/order/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetLogList(ctx context.Context, req *v1.GetLogListReq) (res *v1.GetLogListRes, err error) {
	total, list, err := service.Order().GetLogList(ctx, req.LogQuery)
	if err != nil {
		return nil, err
	}
	res = &v1.GetLogListRes{
		Total: total,
		List:  list,
	}
	return
}
