package permission

import (
	"context"

	"server/internal/dao"
	dao_permission "server/internal/type/permission/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// GetAll implements service.IPermission.
func (s *sPermission) GetAll(ctx context.Context) (res []*dao_permission.List, err error) {
	err = dao.SysPermission.Ctx(ctx).OrderDesc(dao.SysPermission.Columns().CreateTime).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	return
}
