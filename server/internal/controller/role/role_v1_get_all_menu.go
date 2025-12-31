package role

import (
	"context"

	v1 "server/api/role/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetAllMenu(ctx context.Context, req *v1.GetAllMenuReq) (res *v1.GetAllMenuRes, err error) {
	list, err := service.Menu().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetAllMenuRes{
		List: list,
	}
	return
}
