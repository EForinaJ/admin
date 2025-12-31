package user

import (
	"context"

	"server/internal/dao"
	dto_user "server/internal/type/user/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckCreate implements service.IUser.
func (s *sUser) CheckCreate(ctx context.Context, req *dto_user.Create) (err error) {
	res, err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Phone, req.Phone).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if res {
		return utils_error.Err(response.FAILD, "手机号已存在")
	}
	return
}
