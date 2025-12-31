package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/consts"
	"server/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) GetAboutUs(ctx context.Context, req *v1.GetAboutUsReq) (res *v1.GetAboutUsRes, err error) {
	options, err := service.System().GetOne(ctx, consts.AboutUs)
	if err != nil {
		return nil, err
	}
	res = &v1.GetAboutUsRes{
		Content: gconv.String(options),
	}
	return
}
