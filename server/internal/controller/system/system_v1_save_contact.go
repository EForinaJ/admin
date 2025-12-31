package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/consts"
	"server/internal/service"
)

func (c *ControllerV1) SaveContact(ctx context.Context, req *v1.SaveContactReq) (res *v1.SaveContactRes, err error) {
	err = service.System().SaveConfig(ctx, consts.ContactSetting, "客服设置", req.Contact)
	return
}
