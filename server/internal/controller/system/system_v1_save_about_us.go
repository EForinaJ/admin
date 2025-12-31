package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/consts"
	"server/internal/service"
)

func (c *ControllerV1) SaveAboutUs(ctx context.Context, req *v1.SaveAboutUsReq) (res *v1.SaveAboutUsRes, err error) {
	err = service.System().SaveConfig(ctx, consts.AboutUs, "关于我们", req.Content)
	return
}
