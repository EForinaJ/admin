package role

import (
	"context"

	v1 "server/api/role/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetMenus(ctx context.Context, req *v1.GetMenusReq) (res *v1.GetMenusRes, err error) {
	ids, err := service.Menu().GetRelatedId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetMenusRes{
		Menus: ids,
	}
	return
}
