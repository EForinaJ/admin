package witkey

import (
	"context"

	v1 "server/api/witkey/v1"
	"server/internal/consts"
	"server/internal/service"
)

func (c *ControllerV1) ChangeCommission(ctx context.Context, req *v1.ChangeCommissionReq) (res *v1.ChangeCommissionRes, err error) {
	if req.Mode == consts.Sub {
		err := service.Witkey().CheckCommission(ctx, req.Commission)
		if err != nil {
			return nil, err
		}
	}
	err = service.Witkey().ChangeCommission(ctx, req.Commission)
	return
}
