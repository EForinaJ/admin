package user

import (
	"context"

	v1 "server/api/user/v1"
	"server/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = service.User().CheckCreate(ctx, req.Create)
	if err != nil {
		return nil, err
	}

	err = service.User().Create(ctx, req.Create)
	if err != nil {
		return nil, err
	}
	return
}
