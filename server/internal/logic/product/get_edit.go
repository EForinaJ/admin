package product

import (
	"context"
	"server/internal/dao"
	dao_product "server/internal/type/product/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// GetEdit implements service.IProduct.
func (s *sProduct) GetEdit(ctx context.Context, id int64) (res *dao_product.Edit, err error) {
	err = dao.SysProduct.Ctx(ctx).WherePri(id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	return
}
