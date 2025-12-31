package category

import (
	"context"

	"server/internal/dao"
	dto_category "server/internal/type/category/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckEdit implements service.ICategory.
func (s *sCategory) CheckEdit(ctx context.Context, req *dto_category.Edit) (err error) {
	res, err := dao.SysCategory.Ctx(ctx).
		Where(dao.SysCategory.Columns().Name, req.Name).
		WhereNotIn(dao.SysCategory.Columns().Id, req.Id).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if res {
		return utils_error.Err(response.FAILD, "已存在")
	}

	return
}
