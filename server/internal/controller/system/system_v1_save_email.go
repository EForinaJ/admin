package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/consts"
	"server/internal/service"
)

func (c *ControllerV1) SaveEmail(ctx context.Context, req *v1.SaveEmailReq) (res *v1.SaveEmailRes, err error) {
	err = service.System().SaveConfig(ctx, consts.EmailSetting, "邮箱设置", req.Email)
	return
}
