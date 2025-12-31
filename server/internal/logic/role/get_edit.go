package role

import (
	"context"

	"server/internal/dao"
	dao_role "server/internal/type/role/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// GetEdit implements service.IRole.
func (s *sRole) GetEdit(ctx context.Context, id int64) (res *dao_role.Edit, err error) {
	err = dao.SysRole.Ctx(ctx).WherePri(id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	return
}
