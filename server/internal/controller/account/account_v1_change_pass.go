package account

import (
	"context"

	v1 "server/api/account/v1"
	"server/internal/service"
)

func (c *ControllerV1) ChangePass(ctx context.Context, req *v1.ChangePassReq) (res *v1.ChangePassRes, err error) {
	err = service.Account().CheckPass(ctx, req.ChangePass)
	if err != nil {
		return nil, err
	}
	err = service.Account().ChangePass(ctx, req.ChangePass)
	return
}
