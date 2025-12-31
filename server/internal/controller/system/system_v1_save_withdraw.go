package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/consts"
	"server/internal/service"
)

func (c *ControllerV1) SaveWithdraw(ctx context.Context, req *v1.SaveWithdrawReq) (res *v1.SaveWithdrawRes, err error) {
	err = service.System().SaveConfig(ctx, consts.WithdrawSetting, "提现设置", req.Withdraw)
	return
}
