package menu

import (
	"context"

	v1 "server/api/menu/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetEdit(ctx context.Context, req *v1.GetEditReq) (res *v1.GetEditRes, err error) {
	detail, err := service.Menu().GetEdit(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetEditRes{
		Edit: detail,
	}
	return
}
