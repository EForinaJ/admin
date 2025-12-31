package witkey

import (
	"context"

	v1 "server/api/witkey/v1"
	"server/internal/service"
)

func (c *ControllerV1) Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error) {
	err = service.Witkey().Edit(ctx, req.Edit)
	return
}
