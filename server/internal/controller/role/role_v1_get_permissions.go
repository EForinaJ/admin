package role

import (
	"context"

	v1 "server/api/role/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetPermissions(ctx context.Context, req *v1.GetPermissionsReq) (res *v1.GetPermissionsRes, err error) {
	ids, err := service.Permission().GetRelatedId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetPermissionsRes{
		Permissions: ids,
	}
	return
}
