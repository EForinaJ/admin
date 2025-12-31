package permission

import (
	"context"
	"server/internal/dao"
	dao_permission "server/internal/type/permission/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// GetEdit implements service.IPermission.
func (s *sPermission) GetEdit(ctx context.Context, id int64) (res *dao_permission.Edit, err error) {
	err = dao.SysPermission.Ctx(ctx).WherePri(id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	return
}
