package prestore

import (
	"context"

	v1 "server/api/prestore/v1"
	"server/internal/service"
)

func (c *ControllerV1) Apply(ctx context.Context, req *v1.ApplyReq) (res *v1.ApplyRes, err error) {
	err = service.Prestore().CheckApply(ctx, req.Apply)
	if err != nil {
		return nil, err
	}
	err = service.Prestore().Apply(ctx, req.Apply)
	return
}
