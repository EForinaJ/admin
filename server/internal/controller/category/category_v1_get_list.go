package category

import (
	"context"

	v1 "server/api/category/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	total, list, err := service.Category().GetList(ctx, req.Query)
	if err != nil {
		return nil, err
	}
	res = &v1.GetListRes{
		Total: total,
		List:  list,
	}
	return
}
