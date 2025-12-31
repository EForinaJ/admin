package recharge

import (
	"context"
	"server/internal/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// Delete implements service.IRecharge.
func (s *sRecharge) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = dao.SysRecharge.Ctx(ctx).
		WhereIn(dao.SysRecharge.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED, response.CodeMsg(response.DELETE_FAILED))
	}
	return
}
