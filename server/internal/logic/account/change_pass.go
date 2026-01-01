package account

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_account "server/internal/type/account/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

// ChangePass implements service.IAccount.
func (s *sAccount) ChangePass(ctx context.Context, req *dto_account.ChangePass) (err error) {
	newSalt := grand.S(6)
	newToken := consts.SYSTEMNAME + req.ConfirmPass + newSalt
	newToken = gmd5.MustEncryptString(newToken)
	_, err = dao.SysManage.Ctx(ctx).
		Where(dao.SysManage.Columns().Id, ctx.Value("userId")).
		Data(g.Map{
			dao.SysManage.Columns().Salt:       newSalt,
			dao.SysManage.Columns().Password:   newToken,
			dao.SysManage.Columns().UpdateTime: gtime.Now(),
		}).Update()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	return
}
