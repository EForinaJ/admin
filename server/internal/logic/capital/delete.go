package capital

import (
	"context"
	"server/internal/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// Delete implements service.ICapital.
func (s *sCapital) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = dao.SysCapital.Ctx(ctx).
		WhereIn(dao.SysCapital.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR, response.CodeMsg(response.DB_SAVE_ERROR))
	}
	return
}
