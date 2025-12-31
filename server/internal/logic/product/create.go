package product

import (
	"context"
	"server/internal/dao"
	"server/internal/model/do"
	dto_product "server/internal/type/product/dto"
	utils_code "server/internal/utils/code"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Create implements service.IProduct.
func (s *sProduct) Create(ctx context.Context, req *dto_product.Create) (err error) {

	var entity *do.SysProduct
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}

	entity.Code = utils_code.GetCode(ctx, "SN")
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()

	_, err = dao.SysProduct.Ctx(ctx).Data(&entity).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	return
}
