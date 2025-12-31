package title

import (
	"context"
	"server/internal/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// Delete implements service.ITitle.
func (s *sTitle) Delete(ctx context.Context, ids []int64) (err error) {

	_, err = dao.SysTitle.Ctx(ctx).
		WhereIn(dao.SysTitle.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED, response.CodeMsg(response.DELETE_FAILED))
	}
	return
}
