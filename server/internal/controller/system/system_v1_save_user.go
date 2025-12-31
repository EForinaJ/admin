package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/consts"
	"server/internal/service"
)

func (c *ControllerV1) SaveUser(ctx context.Context, req *v1.SaveUserReq) (res *v1.SaveUserRes, err error) {
	err = service.System().SaveConfig(ctx, consts.UserSetting, "用户设置", req.User)
	return
}
