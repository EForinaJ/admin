package witkey

import (
	"context"
	"server/internal/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckExist implements service.IWitkey.
func (s *sWitkey) CheckExist(ctx context.Context, userId int64) (err error) {
	exist, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, userId).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if !exist {
		return utils_error.Err(response.FAILD, "添加的用户不存在")
	}
	exist, err = dao.SysWitkey.Ctx(ctx).Where(dao.SysWitkey.Columns().UserId, userId).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if exist {
		return utils_error.Err(response.FAILD, "该用户已添加威客信息")
	}
	return nil
}
