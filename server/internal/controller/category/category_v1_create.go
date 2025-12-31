package category

import (
	"context"

	v1 "server/api/category/v1"
	"server/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = service.Category().CheckCreate(ctx, req.Create)
	if err != nil {
		return nil, err
	}
	err = service.Category().Create(ctx, req.Create)
	return
}
