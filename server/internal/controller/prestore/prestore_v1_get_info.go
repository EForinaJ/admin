package prestore

import (
	"context"

	v1 "server/api/prestore/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error) {
	detail, err := service.Prestore().GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetInfoRes{
		Detail: detail,
	}
	return
}
