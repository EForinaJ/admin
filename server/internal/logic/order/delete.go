package order

import (
	"context"

	"server/internal/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// Delete implements service.IOrder.
func (s *sOrder) Delete(ctx context.Context, ids []int64) (err error) {

	_, err = dao.SysOrder.Ctx(ctx).
		WhereIn(dao.SysOrder.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED, response.CodeMsg(response.DELETE_FAILED))
	}
	return
}
