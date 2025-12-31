package product

import (
	"context"
	"server/internal/dao"
	"server/internal/model/do"
	dto_product "server/internal/type/product/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Edit implements service.IProduct.
func (s *sProduct) Edit(ctx context.Context, req *dto_product.Edit) (err error) {

	var entity *do.SysProduct
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}

	entity.UpdateTime = gtime.Now()

	_, err = dao.SysProduct.Ctx(ctx).
		Where(dao.SysProduct.Columns().Id, req.Id).
		Data(&entity).Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}

	return
}
