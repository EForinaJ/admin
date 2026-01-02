package order

import (
	"context"

	"server/internal/consts"
	"server/internal/dao"
	dto_order "server/internal/type/order/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// CheckPaid implements service.IOrder.
func (s *sOrder) CheckPaid(ctx context.Context, req *dto_order.Paid) (err error) {
	order, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, req.Id).
		Fields(dao.SysOrder.Columns().Status,
			dao.SysOrder.Columns().UserId,
			dao.SysOrder.Columns().ActualAmount).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusPendingOrder {
		return utils_error.Err(response.FAILD, "订单待服务，无法确认收款")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusInProgress {
		return utils_error.Err(response.FAILD, "订单进行中，无法确认收款")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusComplete {
		return utils_error.Err(response.FAILD, "订单已完成，无法确认收款")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusCancel {
		return utils_error.Err(response.FAILD, "订单已取消，无法确认收款")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusRefund {
		return utils_error.Err(response.FAILD, "订单已退款，无法确认收款")
	}

	if req.PayMode == consts.PayModeBalance {
		userBalance, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Id, order.GMap().Get(dao.SysOrder.Columns().UserId)).
			Value(dao.SysUser.Columns().Balance)
		if err != nil {
			return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		if !decimal.NewFromFloat(userBalance.Float64()).
			GreaterThanOrEqual(decimal.NewFromFloat(gconv.Float64(order.GMap().Get(dao.SysOrder.Columns().ActualAmount)))) {
			return utils_error.Err(response.FAILD, "余额不足")
		}
	}

	return
}
