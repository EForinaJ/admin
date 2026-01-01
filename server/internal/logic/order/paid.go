package order

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	utils_code "server/internal/utils/code"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Paid implements service.IOrder.
func (s *sOrder) Paid(ctx context.Context, id int64) (err error) {
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
	order, err := tx.Model(dao.SysOrder.Table()).
		Where(dao.SysOrder.Columns().Id, id).
		Fields(
			dao.SysOrder.Columns().ActualAmount,
			dao.SysOrder.Columns().UserId,
			dao.SysOrder.Columns().Code,
		).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	_, err = tx.Model(dao.SysOrder.Table()).
		Where(dao.SysOrder.Columns().Id, id).
		Data(g.Map{
			dao.SysOrder.Columns().Status:    consts.OrderStatusPendingOrder,
			dao.SysOrder.Columns().PayStatus: consts.PayStatusSuccess,
			dao.SysOrder.Columns().PayTime:   gtime.Now(),
			dao.SysOrder.Columns().PayMode:   consts.PayModePersonalTransfer,
			dao.SysOrder.Columns().PayAmount: order.GMap().Get(dao.SysOrder.Columns().ActualAmount),
		}).
		Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}

	_, err = tx.Model(dao.SysUserBill.Table()).
		Data(g.Map{
			dao.SysUserBill.Columns().UserId:     order.GMap().Get(dao.SysOrder.Columns().UserId),
			dao.SysUserBill.Columns().RelatedId:  id,
			dao.SysUserBill.Columns().Code:       utils_code.GetCode(ctx, consts.BL),
			dao.SysUserBill.Columns().Type:       consts.BillTypeOrder,
			dao.SysUserBill.Columns().Amount:     order.GMap().Get(dao.SysOrder.Columns().ActualAmount),
			dao.SysUserBill.Columns().Mode:       consts.Sub,
			dao.SysUserBill.Columns().CreateTime: gtime.Now(),
		}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	//  添加支付日志
	_, err = tx.Model(dao.SysCapital.Table()).Data(g.Map{
		dao.SysCapital.Columns().CreateTime: gtime.Now(),
		dao.SysCapital.Columns().Code:       utils_code.GetCode(ctx, consts.PM),
		dao.SysCapital.Columns().Related:    order.GMap().Get(dao.SysOrder.Columns().Code),
		dao.SysCapital.Columns().Amount:     order.GMap().Get(dao.SysOrder.Columns().ActualAmount),
		dao.SysCapital.Columns().Type:       consts.CapitalPaymentOrder,
		dao.SysCapital.Columns().Mode:       consts.PayModePersonalTransfer,
		dao.SysCapital.Columns().UserId:     order.GMap().Get(dao.SysOrder.Columns().UserId),
	}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	//  添加订单日志
	_, err = tx.Model(dao.SysOrderLog.Table()).Data(g.Map{
		dao.SysOrderLog.Columns().CreateTime: gtime.Now(),
		dao.SysOrderLog.Columns().Content:    "订单手动收款",
		dao.SysOrderLog.Columns().ManageId:   ctx.Value("userId"),
		dao.SysOrderLog.Columns().OrderId:    id,
		dao.SysOrderLog.Columns().Type:       consts.OrderLogTypePaid,
	}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	return
}
