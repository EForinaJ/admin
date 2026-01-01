package order

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_order "server/internal/type/order/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// CheckDistribute implements service.IOrder.
func (s *sOrder) CheckDistribute(ctx context.Context, req *dto_order.Distribute) (err error) {
	// 判断威客是否存在
	witkeyStatus, err := dao.SysWitkey.Ctx(ctx).Where(dao.SysWitkey.Columns().Id, req.WitkeyId).
		Value(dao.SysWitkey.Columns().Status)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	// 判断是否在线
	if witkeyStatus.Int() != consts.Enable {
		return utils_error.Err(response.FAILD, "威客禁用中")
	}
	// 判断是否已经派发过
	distributeExist, err := dao.SysDistribute.Ctx(ctx).
		Where(dao.SysDistribute.Columns().OrderId, req.Id).
		Where(dao.SysDistribute.Columns().WitkeyId, req.WitkeyId).
		Where(dao.SysDistribute.Columns().IsCancel, consts.Not).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if distributeExist {
		return utils_error.Err(response.FAILD, "已派发该威客，请勿重复派发")
	}

	// 有售后工单无法派单
	aftersalesExist, err := dao.SysAftersales.Ctx(ctx).Where(dao.SysAftersales.Columns().OrderId, req.Id).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if aftersalesExist {
		return utils_error.Err(response.FAILD, "该订单提交售后工单，无法重新派单")
	}

	order, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, req.Id).
		Fields(dao.SysOrder.Columns().WitkeyCount, dao.SysOrder.Columns().Status).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusPendingPayment {
		return utils_error.Err(response.FAILD, "订单待支付，无法开始")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusComplete {
		return utils_error.Err(response.FAILD, "订单已完成，无法开始")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusCancel {
		return utils_error.Err(response.FAILD, "订单已取消，无法开始")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusRefund {
		return utils_error.Err(response.FAILD, "订单已退款，无法开始")
	}

	count, err := dao.SysDistribute.Ctx(ctx).
		Where(dao.SysDistribute.Columns().OrderId, req.Id).
		Where(dao.SysDistribute.Columns().IsCancel, consts.Not).Count()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if count >= gconv.Int(order.GMap().Get(dao.SysOrder.Columns().WitkeyCount)) {
		return utils_error.Err(response.FAILD, "服务人数已上限")
	}
	return
}
