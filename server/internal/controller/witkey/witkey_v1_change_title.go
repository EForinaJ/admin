package witkey

import (
	"context"

	v1 "server/api/witkey/v1"
	"server/internal/service"
)

func (c *ControllerV1) ChangeTitle(ctx context.Context, req *v1.ChangeTitleReq) (res *v1.ChangeTitleRes, err error) {
	err = service.Witkey().CheckChangeTitle(ctx, req.ChangeTitle)
	if err != nil {
		return nil, err
	}

	err = service.Witkey().ChangeTitle(ctx, req.ChangeTitle)
	return
}
