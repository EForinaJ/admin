package aftersales

import (
	"context"

	v1 "server/api/aftersales/v1"
	"server/internal/service"
)

func (c *ControllerV1) Apply(ctx context.Context, req *v1.ApplyReq) (res *v1.ApplyRes, err error) {
	err = service.Aftersales().CheckApply(ctx, req.Apply)
	if err != nil {
		return nil, err
	}
	err = service.Aftersales().Apply(ctx, req.Apply)
	return
}
