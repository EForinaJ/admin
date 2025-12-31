package role

import (
	"context"
	"server/internal/dao"
	dto_role "server/internal/type/role/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckEdit implements service.IRole.
func (s *sRole) CheckEdit(ctx context.Context, req *dto_role.Edit) (err error) {
	exist, err := dao.SysRole.Ctx(ctx).
		Where(dao.SysRole.Columns().Code, req.Code).
		WhereNotIn(dao.SysRole.Columns().Id, req.Id).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if exist {
		return utils_error.Err(response.FAILD, "角色已存在")
	}

	return
}
