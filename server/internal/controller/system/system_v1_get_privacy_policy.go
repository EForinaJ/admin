package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/consts"
	"server/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) GetPrivacyPolicy(ctx context.Context, req *v1.GetPrivacyPolicyReq) (res *v1.GetPrivacyPolicyRes, err error) {
	options, err := service.System().GetOne(ctx, consts.PrivacyPolicy)
	if err != nil {
		return nil, err
	}
	res = &v1.GetPrivacyPolicyRes{
		Content: gconv.String(options),
	}
	return
}
