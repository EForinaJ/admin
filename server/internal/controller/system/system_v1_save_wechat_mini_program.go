package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/consts"
	"server/internal/service"
)

func (c *ControllerV1) SaveWechatMiniProgram(ctx context.Context, req *v1.SaveWechatMiniProgramReq) (res *v1.SaveWechatMiniProgramRes, err error) {
	err = service.System().SaveConfig(ctx, consts.WechatMiniProgramSetting, "微信小程序设置", req.WechatMiniProgram)
	return
}
