package comment

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_comment "server/internal/type/cooment/dao"
	dto_comment "server/internal/type/cooment/dto"

	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IComment.
func (s *sComment) GetList(ctx context.Context, req *dto_comment.Query) (total int, res []*dao_comment.List, err error) {
	m := dao.SysComment.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysComment.Columns().CreateTime)

	if req.Code != "" {
		productId, err := dao.SysProduct.Ctx(ctx).
			Where(dao.SysProduct.Columns().Code, req.Code).
			Value(dao.SysProduct.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		m = m.Where(dao.SysComment.Columns().ProductId, productId)
	}
	if req.Status != 0 {
		m = m.Where(dao.SysComment.Columns().Status, req.Status)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var list []*entity.SysComment
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_comment.List, len(list))
	for i, v := range list {
		var entity *dao_comment.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		// 评论用户
		user, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Id, v.UserId).
			Value(dao.SysUser.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.User = user.String()

		// 商品内容
		product, err := dao.SysProduct.Ctx(ctx).
			Where(dao.SysProduct.Columns().Id, v.ProductId).
			Fields(
				dao.SysProduct.Columns().Name,
				dao.SysProduct.Columns().Pic,
				dao.SysProduct.Columns().CategoryId,
			).
			One()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		category, err := dao.SysCategory.Ctx(ctx).
			Where(dao.SysCategory.Columns().Id, product.GMap().Get(dao.SysProduct.Columns().CategoryId)).
			Value(dao.SysCategory.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		entity.Product = &dao_comment.Product{
			Id:       v.ProductId,
			Name:     gconv.String(product.GMap().Get(dao.SysProduct.Columns().Name)),
			Pic:      gconv.String(product.GMap().Get(dao.SysProduct.Columns().Pic)),
			Category: category.String(),
		}
		res[i] = entity
	}

	return
}
