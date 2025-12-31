package product

import (
	"context"
	"server/internal/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// Delete implements service.IProduct.
func (s *sProduct) Delete(ctx context.Context, ids []int64) (err error) {
	// skus, err := tx.Model(dao.SysProductSku.Table()).
	// 	Fields(dao.SysProductSku.Columns().Id).
	// 	WhereIn(dao.SysProductSku.Columns().ProductId, ids).Array()
	// if err != nil {
	// 	return utils_error.Err(response.DB_READ_ERROR)
	// }

	// _, err = tx.Model(dao.SysProductSkuSpecRelations.Table()).
	// 	WhereIn(dao.SysProductSkuSpecRelations.Columns().SkuId, skus).Delete()
	// if err != nil {
	// 	return utils_error.Err(response.DB_READ_ERROR)
	// }

	// attrs, err := tx.Model(dao.SysProductSpecAttrs.Table()).
	// 	Fields(dao.SysProductSpecAttrs.Columns().Id).
	// 	WhereIn(dao.SysProductSpecAttrs.Columns().ProductId, ids).Array()
	// if err != nil {
	// 	return utils_error.Err(response.DB_READ_ERROR)
	// }

	// _, err = tx.Model(dao.SysProductSpecValues.Table()).
	// 	WhereIn(dao.SysProductSpecValues.Columns().SpecAttsId, attrs).Delete()
	// if err != nil {
	// 	return utils_error.Err(response.DB_READ_ERROR)
	// }

	// _, err = tx.Model(dao.SysProductSpecAttrs.Table()).
	// 	WhereIn(dao.SysProductSpecAttrs.Columns().ProductId, ids).Delete()
	// if err != nil {
	// 	return utils_error.Err(response.DB_READ_ERROR)
	// }

	// _, err = tx.Model(dao.SysProductSku.Table()).
	// 	WhereIn(dao.SysProductSku.Columns().ProductId, ids).Delete()
	// if err != nil {
	// 	return utils_error.Err(response.DB_READ_ERROR)
	// }

	_, err = dao.SysProduct.Ctx(ctx).
		WhereIn(dao.SysProduct.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED, response.CodeMsg(response.DELETE_FAILED))
	}

	return
}
