package witkey

import (
	"context"

	v1 "server/api/witkey/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetCommissionList(ctx context.Context, req *v1.GetCommissionListReq) (res *v1.GetCommissionListRes, err error) {
	total, list, err := service.Witkey().GetCommissionList(ctx, req.CommissionQuery)
	if err != nil {
		return nil, err
	}
	res = &v1.GetCommissionListRes{
		Total: total,
		List:  list,
	}

	return
}
