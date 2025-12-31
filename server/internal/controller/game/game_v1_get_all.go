package game

import (
	"context"

	v1 "server/api/game/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetAll(ctx context.Context, req *v1.GetAllReq) (res *v1.GetAllRes, err error) {
	list, err := service.Game().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetAllRes{
		List: list,
	}
	return
}
