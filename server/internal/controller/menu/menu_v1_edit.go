package menu

import (
	"context"

	v1 "server/api/menu/v1"
	"server/internal/service"
)

func (c *ControllerV1) Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error) {
	err = service.Menu().CheckEdit(ctx, req.Edit)
	if err != nil {
		return nil, err
	}

	err = service.Menu().Edit(ctx, req.Edit)
	if err != nil {
		return nil, err
	}
	return
}
