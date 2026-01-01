package account

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_account "server/internal/type/account/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// Edit implements service.IAccount.
func (s *sAccount) Edit(ctx context.Context, req *dto_account.Edit) (err error) {
	_, err = dao.SysManage.Ctx(ctx).
		Where(dao.SysManage.Columns().Id, ctx.Value("userId")).
		Data(g.Map{
			dao.SysManage.Columns().Avatar:      req.Avatar,
			dao.SysManage.Columns().Name:        req.Name,
			dao.SysManage.Columns().Email:       req.Email,
			dao.SysManage.Columns().Address:     req.Address,
			dao.SysManage.Columns().Sex:         req.Sex,
			dao.SysManage.Columns().Description: req.Description,
			dao.SysManage.Columns().Tags:        req.Tags,
		}).Update()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	_, err = g.Redis().Del(ctx, consts.Account+gconv.String(ctx.Value("userId")))
	if err != nil {
		return utils_error.Err(response.CACHE_SAVE_ERROR, response.CodeMsg(response.CACHE_SAVE_ERROR))
	}
	return
}
