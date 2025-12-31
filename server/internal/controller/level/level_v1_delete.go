package level

import (
	"context"

	v1 "server/api/level/v1"
	"server/internal/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.Level().Delete(ctx, req.Ids)
	return
}
