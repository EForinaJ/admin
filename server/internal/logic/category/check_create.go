package category

import (
	"context"

	"server/internal/dao"
	dto_category "server/internal/type/category/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckCreate implements service.ICategory.
func (s *sCategory) CheckCreate(ctx context.Context, req *dto_category.Create) (err error) {
	res, err := dao.SysCategory.Ctx(ctx).
		Where(dao.SysCategory.Columns().Name, req.Name).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if res {
		return utils_error.Err(response.FAILD, "已存在")
	}
	return
}
