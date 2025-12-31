package category

import (
	"context"
	"server/internal/dao"
	dao_category "server/internal/type/category/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// GetEdit implements service.ICategory.
func (s *sCategory) GetEdit(ctx context.Context, id int64) (res *dao_category.Edit, err error) {
	err = dao.SysCategory.Ctx(ctx).Where(dao.SysCategory.Columns().Id, id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	return
}
