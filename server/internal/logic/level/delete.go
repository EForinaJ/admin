package level

import (
	"context"

	"server/internal/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// Delete implements service.ILevel.
func (s *sLevel) Delete(ctx context.Context, ids []int64) (err error) {

	_, err = dao.SysLevel.Ctx(ctx).
		WhereIn(dao.SysLevel.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED, response.CodeMsg(response.DELETE_FAILED))
	}
	return
}
