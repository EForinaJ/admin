package account

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dao_account "server/internal/type/account/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IAccount.
func (s *sAccount) GetDetail(ctx context.Context) (res *dao_account.Detail, err error) {
	userId := ctx.Value("userId")
	options, err := g.Redis().Get(ctx, consts.Account+gconv.String(userId))
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
	}

	if !options.IsEmpty() {
		err = options.Scan(&res)
		if err != nil {
			return nil, utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
		}
		return
	}

	err = dao.SysManage.Ctx(ctx).Fields(dao.SysManage.Columns().Id,
		dao.SysManage.Columns().Name,
		dao.SysManage.Columns().Email,
		dao.SysManage.Columns().Address,
		dao.SysManage.Columns().Description,
		dao.SysManage.Columns().Sex,
		dao.SysManage.Columns().Tags,
		dao.SysManage.Columns().Id,
		dao.SysManage.Columns().Avatar).
		Where(dao.SysManage.Columns().Id, userId).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if res.Tags == nil {
		res.Tags = []string{}
	}
	// 获取角色
	rolesList, err := dao.SysManageRole.
		Ctx(ctx).
		Where(dao.SysManageRole.Columns().ManageId, userId).
		Fields(dao.SysManageRole.Columns().RoleId).
		Array()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	roles, err := dao.SysRole.
		Ctx(ctx).
		Where(dao.SysRole.Columns().Id, rolesList).
		Fields(dao.SysRole.Columns().Name).
		Array()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	res.Roles = roles.Strings()

	err = g.Redis().SetEX(ctx, consts.Account+gconv.String(userId), res, 600)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_SAVE_ERROR, response.CodeMsg(response.CACHE_SAVE_ERROR))
	}
	return
}
