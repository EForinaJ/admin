package site

import (
	"context"

	v1 "server/api/site/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetUserOptions(ctx context.Context, req *v1.GetUserOptionsReq) (res *v1.GetUserOptionsRes, err error) {
	list, err := service.Site().GetUserOptions(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	res = &v1.GetUserOptionsRes{
		List: list,
	}
	return
}
