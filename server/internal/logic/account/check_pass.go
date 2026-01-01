package account

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_account "server/internal/type/account/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
	"strings"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
)

// CheckPass implements service.IAccount.
func (s *sAccount) CheckPass(ctx context.Context, req *dto_account.ChangePass) (err error) {
	manage, err := dao.SysManage.Ctx(ctx).
		Where(dao.SysManage.Columns().Id, ctx.Value("userId")).
		Fields(dao.SysManage.Columns().Password, dao.SysManage.Columns().Salt).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	salt := gconv.String(manage.GMap().Get(dao.SysManage.Columns().Salt))

	randPwd := consts.SYSTEMNAME + req.OldPass + salt
	randPwd = gmd5.MustEncryptString(randPwd)
	if !strings.EqualFold(gconv.String(manage.GMap().Get(dao.SysManage.Columns().Password)), randPwd) {
		return utils_error.Err(response.UPDATE_FAILED, "密码不正确")
	}

	return
}
