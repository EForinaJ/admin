package permission

import (
	"context"

	"server/internal/dao"
	dto_permission "server/internal/type/permission/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckEdit implements service.IPermission.
func (s *sPermission) CheckEdit(ctx context.Context, req *dto_permission.Edit) (err error) {
	exist, err := dao.SysPermission.Ctx(ctx).
		Where(dao.SysPermission.Columns().Permission, req.Permission).
		WhereNotIn(dao.SysPermission.Columns().Id, req.Id).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if exist {
		return utils_error.Err(response.FAILD, "权限标识已存在")
	}
	return
}
