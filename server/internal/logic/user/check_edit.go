package user

import (
	"context"

	"server/internal/dao"
	dto_user "server/internal/type/user/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckEdit implements service.IUser.
func (s *sUser) CheckEdit(ctx context.Context, req *dto_user.Edit) (err error) {
	res, err := dao.SysUser.Ctx(ctx).
		WhereNotIn(dao.SysUser.Columns().Id, req.Id).
		Where(dao.SysUser.Columns().Phone, req.Phone).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if res {
		return utils_error.Err(response.FAILD, "手机号已存在")
	}
	return
}
