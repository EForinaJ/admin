package category

import (
	"context"

	"server/internal/dao"
	"server/internal/model/do"
	dto_category "server/internal/type/category/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Create implements service.ICategory.
func (s *sCategory) Create(ctx context.Context, req *dto_category.Create) (err error) {

	var entity *do.SysCategory
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysCategory.Ctx(ctx).Data(&entity).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR, response.CodeMsg(response.DB_SAVE_ERROR))
	}

	return
}
