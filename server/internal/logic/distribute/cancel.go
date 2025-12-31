package distribute

import (
	"context"

	"server/internal/consts"
	"server/internal/dao"
	dto_distribute "server/internal/type/distribute/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
)

// Cancel implements service.IDistribute.
func (s *sDistribute) Cancel(ctx context.Context, req *dto_distribute.Cancel) (err error) {
	_, err = dao.SysDistribute.Ctx(ctx).Where(dao.SysDistribute.Columns().Id, req.Id).Data(g.Map{
		dao.SysDistribute.Columns().IsCancel: consts.Yes,
		dao.SysDistribute.Columns().Reason:   req.Reason,
	}).Update()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	return
}
