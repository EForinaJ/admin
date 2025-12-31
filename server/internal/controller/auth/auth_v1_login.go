package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "server/api/auth/v1"
	"server/internal/consts"
	"server/internal/service"
	utlis_lock "server/internal/utils/lock"
	"server/internal/utils/response"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	//检查是否用户被锁
	if utlis_lock.CheckLock(ctx, consts.LoginLock+req.Phone) {
		return nil, gerror.NewCode(gcode.New(gconv.Int(response.LOGIN_ERROR), "", nil), "账号已锁定，请30分钟后再试")
	}

	// 获取token
	token, err := service.Auth().Login(ctx, req.Login)
	if err != nil {
		return nil, err
	}
	res = &v1.LoginRes{
		Token: token.(string),
	}
	return
}
