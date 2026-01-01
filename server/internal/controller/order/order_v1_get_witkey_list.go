package order

import (
	"context"

	v1 "server/api/order/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetWitkeyList(ctx context.Context, req *v1.GetWitkeyListReq) (res *v1.GetWitkeyListRes, err error) {
	total, list, err := service.Witkey().GetList(ctx, req.Query)
	if err != nil {
		return nil, err
	}
	res = &v1.GetWitkeyListRes{
		Total: total,
		List:  list,
	}
	return
}
