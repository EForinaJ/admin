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

// CheckDiscount implements service.IOrder.
func (s *sOrder) CheckDiscount(ctx context.Context, req *dto_order.AddDiscount) (err error) {
	order, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, req.Id).
		Fields(dao.SysOrder.Columns().TotalAmount, dao.SysOrder.Columns().Status).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if !decimal.NewFromFloat(req.Amount).LessThanOrEqual(decimal.NewFromFloat(gconv.Float64(order.GMap().Get(dao.SysOrder.Columns().TotalAmount)))) {
		return utils_error.Err(response.FAILD, "优惠金额超出订单总额")
	}

	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusPendingOrder {
		return utils_error.Err(response.FAILD, "订单待服务，无法添加优惠")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusInProgress {
		return utils_error.Err(response.FAILD, "订单进行中，无法添加优惠")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusComplete {
		return utils_error.Err(response.FAILD, "订单已完成，无法添加优惠")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusCancel {
		return utils_error.Err(response.FAILD, "订单已取消，无法添加优惠")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusRefund {
		return utils_error.Err(response.FAILD, "订单已退款，无法添加优惠")
	}

	return nil
}
