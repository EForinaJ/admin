package menu

import (
	"context"

	v1 "server/api/menu/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetAll(ctx context.Context, req *v1.GetAllReq) (res *v1.GetAllRes, err error) {
	list, err := service.Menu().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetAllRes{
		List: list,
	}
	return
}
