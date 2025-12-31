package manage

import (
	"context"

	"server/internal/dao"
	dao_manage "server/internal/type/manage/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetEdit implements service.IManage.
func (s *sManage) GetEdit(ctx context.Context, id int64) (res *dao_manage.Edit, err error) {
	err = dao.SysManage.Ctx(ctx).WherePri(id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	// 获取管理员的角色
	roleIds, err := dao.SysManageRole.Ctx(ctx).
		Fields(dao.SysManageRole.Columns().RoleId).
		Where(dao.SysManageRole.Columns().ManageId, id).
		Array()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	res.Roles = gconv.Int64s(roleIds)

	return
}
