package role

import (
	"context"

	v1 "server/api/role/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetAllPermission(ctx context.Context, req *v1.GetAllPermissionReq) (res *v1.GetAllPermissionRes, err error) {
	list, err := service.Permission().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetAllPermissionRes{
		List: list,
	}
	return
}
