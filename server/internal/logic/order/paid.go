package order

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_order "server/internal/type/order/dto"
	utils_code "server/internal/utils/code"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// Paid implements service.IOrder.
func (s *sOrder) Paid(ctx context.Context, req *dto_order.Paid) (err error) {
	if req.PayMode == consts.PayModeBalance {
		err = s.balancePaid(ctx, req)
		if err != nil {
			return err
		}
	}
	if req.PayMode == consts.PayModePersonalTransfer {
		err = s.personalTransferPaid(ctx, req)
		if err != nil {
			return err
		}
	}
	return
}

func (s *sOrder) personalTransferPaid(ctx context.Context, req *dto_order.Paid) (err error) {
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
		Where(dao.SysOrder.Columns().Id, req.Id).
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
		Where(dao.SysOrder.Columns().Id, req.Id).
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
			dao.SysUserBill.Columns().RelatedId:  req.Id,
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
		dao.SysOrderLog.Columns().OrderId:    req.Id,
		dao.SysOrderLog.Columns().Type:       consts.OrderLogTypePaid,
	}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	return
}

func (s *sOrder) balancePaid(ctx context.Context, req *dto_order.Paid) (err error) {
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
		Where(dao.SysOrder.Columns().Id, req.Id).
		Fields(
			dao.SysOrder.Columns().ActualAmount,
			dao.SysOrder.Columns().UserId,
			dao.SysOrder.Columns().Code,
		).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	userBalance, err := tx.Model(dao.SysUser.Table()).
		Where(dao.SysUser.Columns().Id, order.GMap().Get(dao.SysOrder.Columns().UserId)).
		Value(dao.SysUser.Columns().Balance)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	balance := decimal.NewFromFloat(userBalance.Float64())
	amount := decimal.NewFromFloat(gconv.Float64(order.GMap().Get(dao.SysOrder.Columns().ActualAmount)))
	newBalance := balance.Sub(amount)
	_, err = tx.Model(dao.SysUser.Table()).
		Where(dao.SysUser.Columns().Id, order.GMap().Get(dao.SysOrder.Columns().UserId)).
		Data(g.Map{
			dao.SysUser.Columns().Balance: newBalance,
		}).Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}

	//  添加日志
	_, err = tx.Model(dao.SysBalance.Table()).Data(g.Map{
		dao.SysBalance.Columns().After:      balance,
		dao.SysBalance.Columns().Amount:     amount,
		dao.SysBalance.Columns().Before:     newBalance,
		dao.SysBalance.Columns().Mode:       consts.Sub,
		dao.SysBalance.Columns().UserId:     order.GMap().Get(dao.SysOrder.Columns().UserId),
		dao.SysBalance.Columns().CreateTime: gtime.Now(),
		dao.SysBalance.Columns().Type:       consts.UserChangeBalanceTypePaidOrder,
		dao.SysBalance.Columns().Remark:     "订单消费",
		dao.SysBalance.Columns().Related:    order.GMap().Get(dao.SysOrder.Columns().Code),
	}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	_, err = tx.Model(dao.SysOrder.Table()).
		Where(dao.SysOrder.Columns().Id, req.Id).
		Data(g.Map{
			dao.SysOrder.Columns().Status:    consts.OrderStatusPendingOrder,
			dao.SysOrder.Columns().PayStatus: consts.PayStatusSuccess,
			dao.SysOrder.Columns().PayTime:   gtime.Now(),
			dao.SysOrder.Columns().PayMode:   consts.PayModeBalance,
			dao.SysOrder.Columns().PayAmount: order.GMap().Get(dao.SysOrder.Columns().ActualAmount),
		}).
		Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}
	_, err = tx.Model(dao.SysUserBill.Table()).
		Data(g.Map{
			dao.SysUserBill.Columns().UserId:     order.GMap().Get(dao.SysOrder.Columns().UserId),
			dao.SysUserBill.Columns().RelatedId:  req.Id,
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
		dao.SysCapital.Columns().Mode:       consts.PayModeBalance,
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
		dao.SysOrderLog.Columns().OrderId:    req.Id,
		dao.SysOrderLog.Columns().Type:       consts.OrderLogTypePaid,
	}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	return
}
