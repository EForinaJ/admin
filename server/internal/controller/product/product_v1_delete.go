package product

import (
	"context"

	v1 "server/api/product/v1"
	"server/internal/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.Product().Delete(ctx, req.Ids)
	return
}
