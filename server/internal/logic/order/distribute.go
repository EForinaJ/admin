package order

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_order "server/internal/type/order/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Distribute implements service.IOrder.
func (s *sOrder) Distribute(ctx context.Context, req *dto_order.Distribute) (err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return utils_error.Err(response.DB_TX_ERROR, response.CodeMsg(response.DB_TX_ERROR))
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Model(dao.SysDistribute.Table()).
		Data(g.Map{
			dao.SysDistribute.Columns().WitkeyId:   req.WitkeyId,
			dao.SysDistribute.Columns().OrderId:    req.Id,
			dao.SysDistribute.Columns().IsCancel:   consts.Not,
			dao.SysDistribute.Columns().ManageId:   ctx.Value("userId"),
			dao.SysDistribute.Columns().CreateTime: gtime.Now(),
		}).
		Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR, response.CodeMsg(response.DB_SAVE_ERROR))
	}

	//  添加订单日志
	_, err = tx.Model(dao.SysOrderLog.Table()).Data(g.Map{
		dao.SysOrderLog.Columns().CreateTime: gtime.Now(),
		dao.SysOrderLog.Columns().Content:    "派发威客",
		dao.SysOrderLog.Columns().ManageId:   ctx.Value("userId"),
		dao.SysOrderLog.Columns().OrderId:    req.Id,
		dao.SysOrderLog.Columns().Type:       consts.OrderLogTypeDistribute,
	}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	return
}
