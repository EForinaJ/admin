package manage

import (
	"context"

	"server/internal/dao"
	dto_manage "server/internal/type/manage/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckCreate implements service.IManage.
func (s *sManage) CheckCreate(ctx context.Context, req *dto_manage.Create) (err error) {
	eixst, err := dao.SysManage.Ctx(ctx).
		Where(dao.SysManage.Columns().Phone, req.Phone).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if eixst {
		return utils_error.Err(response.FAILD, "手机号已存在")
	}
	return
}
