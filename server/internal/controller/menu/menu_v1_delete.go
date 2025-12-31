package menu

import (
	"context"

	v1 "server/api/menu/v1"
	"server/internal/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.Menu().Delete(ctx, req.Ids)
	return
}
