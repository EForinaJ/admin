package product

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_product "server/internal/type/product/dao"
	dto_product "server/internal/type/product/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IProduct.
func (s *sProduct) GetList(ctx context.Context, req *dto_product.Query) (total int, res []*dao_product.List, err error) {
	m := dao.SysProduct.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysProduct.Columns().CreateTime)

	if req.Name != "" {
		m = m.Where(dao.SysProduct.Columns().Name, req.Name)
	}
	if req.GameId != 0 {
		m = m.Where(dao.SysProduct.Columns().GameId, req.GameId)
	}
	if req.Status != 0 {
		m = m.Where(dao.SysProduct.Columns().Status, req.Status)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var list []*entity.SysProduct
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_product.List, len(list))
	for i, v := range list {
		var entity *dao_product.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		game, err := dao.SysGame.Ctx(ctx).
			WherePri(v.GameId).
			Value(dao.SysGame.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Game = game.String()

		category, err := dao.SysCategory.Ctx(ctx).
			WherePri(v.CategoryId).
			Value(dao.SysCategory.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Category = category.String()

		res[i] = entity
	}

	return
}
