package distribute

import (
	"context"

	"server/internal/consts"
	"server/internal/dao"
	dto_distribute "server/internal/type/distribute/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Create implements service.IDistribute.
func (s *sDistribute) Create(ctx context.Context, req *dto_distribute.Create) (err error) {

	order, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Code, req.Code).Value(dao.SysOrder.Columns().Id)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR, response.CodeMsg(response.DB_SAVE_ERROR))
	}
	_, err = dao.SysDistribute.Ctx(ctx).
		Data(g.Map{
			dao.SysDistribute.Columns().WitkeyId:   req.WitkeyId,
			dao.SysDistribute.Columns().OrderId:    order.Int64(),
			dao.SysDistribute.Columns().IsCancel:   consts.Not,
			dao.SysDistribute.Columns().ManageId:   ctx.Value("userId"),
			dao.SysDistribute.Columns().CreateTime: gtime.Now(),
		}).
		Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR, response.CodeMsg(response.DB_SAVE_ERROR))
	}
	return
}
