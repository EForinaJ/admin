package settlement

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_settlement "server/internal/type/settlement/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// CheckApply implements service.ISettlement.
func (s *sSettlement) CheckApply(ctx context.Context, req *dto_settlement.Apply) (err error) {
	obj, err := dao.SysSettlement.Ctx(ctx).Where(dao.SysSettlement.Columns().Id, req.Id).
		Fields(dao.SysSettlement.Columns().OrderId, dao.SysSettlement.Columns().Status).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	orderStatus, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, obj.GMap().Get(dao.SysSettlement.Columns().OrderId)).
		Value(dao.SysOrder.Columns().Status)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if orderStatus.Int() == consts.OrderStatusInProgress {
		return utils_error.Err(response.FAILD, "订单正在进行中，无法审核结算")
	}
	if orderStatus.Int() == consts.OrderStatusPendingOrder {
		return utils_error.Err(response.FAILD, "订单等待服务中，无法审核结算")
	}
	if orderStatus.Int() == consts.OrderStatusCancel {
		return utils_error.Err(response.FAILD, "订单已取消，无法审核结算")
	}
	if orderStatus.Int() == consts.OrderStatusPendingPayment {
		return utils_error.Err(response.FAILD, "订单未支付，无法审核结算")
	}

	if gconv.Int(obj.GMap().Get(dao.SysSettlement.Columns().Status)) != consts.StatusApply {
		return utils_error.Err(response.FAILD, "该报单结算已审核")
	}

	aftersalesStatus, err := dao.SysAftersales.Ctx(ctx).
		Where(dao.SysAftersales.Columns().OrderId, obj.GMap().Get(dao.SysSettlement.Columns().OrderId)).
		Value(dao.SysAftersales.Columns().Status)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if aftersalesStatus.Int() == consts.StatusApply {
		return utils_error.Err(response.FAILD, "订单提交售后工单未处理，无法审核结算")
	}

	return
}
