package order

import (
	"context"

	"server/internal/consts"
	"server/internal/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckPaid implements service.IOrder.
func (s *sOrder) CheckPaid(ctx context.Context, id int64) (err error) {
	status, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, id).
		Value(dao.SysOrder.Columns().Status)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if status.Int() == consts.OrderStatusPendingOrder {
		return utils_error.Err(response.FAILD, "订单待服务，无法确认收款")
	}
	if status.Int() == consts.OrderStatusInProgress {
		return utils_error.Err(response.FAILD, "订单进行中，无法确认收款")
	}
	if status.Int() == consts.OrderStatusComplete {
		return utils_error.Err(response.FAILD, "订单已完成，无法确认收款")
	}
	if status.Int() == consts.OrderStatusCancel {
		return utils_error.Err(response.FAILD, "订单已取消，无法确认收款")
	}
	if status.Int() == consts.OrderStatusRefund {
		return utils_error.Err(response.FAILD, "订单已退款，无法确认收款")
	}

	return
}
