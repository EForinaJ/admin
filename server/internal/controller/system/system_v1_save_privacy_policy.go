package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/consts"
	"server/internal/service"
)

func (c *ControllerV1) SavePrivacyPolicy(ctx context.Context, req *v1.SavePrivacyPolicyReq) (res *v1.SavePrivacyPolicyRes, err error) {
	err = service.System().SaveConfig(ctx, consts.PrivacyPolicy, "隐私协议", req.Content)
	return
}
