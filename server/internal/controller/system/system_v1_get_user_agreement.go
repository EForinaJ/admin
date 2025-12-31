package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/consts"
	"server/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) GetUserAgreement(ctx context.Context, req *v1.GetUserAgreementReq) (res *v1.GetUserAgreementRes, err error) {
	options, err := service.System().GetOne(ctx, consts.UserAgreement)
	if err != nil {
		return nil, err
	}
	res = &v1.GetUserAgreementRes{
		Content: gconv.String(options),
	}
	return
}
