package manage

import (
	"context"

	"server/internal/dao"
	dto_manage "server/internal/type/manage/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckEdit implements service.IManage.
func (s *sManage) CheckEdit(ctx context.Context, req *dto_manage.Edit) (err error) {
	exist, err := dao.SysManage.Ctx(ctx).
		Where(dao.SysManage.Columns().Phone, req.Phone).
		WhereNotIn(dao.SysManage.Columns().Id, req.Id).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if exist {
		return utils_error.Err(response.FAILD, "手机号已存在")
	}

	return
}
