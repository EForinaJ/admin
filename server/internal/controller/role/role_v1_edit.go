package role

import (
	"context"

	v1 "server/api/role/v1"
	"server/internal/service"
)

func (c *ControllerV1) Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error) {
	// 判断是否存在
	err = service.Role().CheckEdit(ctx, req.Edit)
	if err != nil {
		return nil, err
	}

	err = service.Role().Edit(ctx, req.Edit)
	if err != nil {
		return nil, err
	}
	return
}
