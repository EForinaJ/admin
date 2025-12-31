package order

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
)

// Complete implements service.IOrder.
func (s *sOrder) Complete(ctx context.Context, id int64) (err error) {
	_, err = dao.SysOrder.Ctx(ctx).Where(dao.SysOrder.Columns().Id, id).Data(g.Map{
		dao.SysOrder.Columns().Status: consts.OrderStatusComplete,
	}).Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}
	return
}
