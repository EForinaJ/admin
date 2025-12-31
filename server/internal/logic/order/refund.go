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
)

// Refund implements service.IOrder.
func (s *sOrder) Refund(ctx context.Context, req *dto_order.Refund) (err error) {
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
	order, err := tx.Model(dao.SysOrder.Table()).WherePri(req.Id).
		Fields(dao.SysOrder.Columns().PayMode,
			dao.SysOrder.Columns().UserId,
			dao.SysOrder.Columns().Code).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	_, err = tx.Model(dao.SysOrder.Table()).WherePri(req.Id).Data(g.Map{
		dao.SysOrder.Columns().PayStatus: consts.PayStatusRefund,
		dao.SysOrder.Columns().Status:    consts.OrderStatusRefund,
	}).Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}

	_, err = tx.Model(dao.SysAftersales.Table()).
		Data(g.Map{
			dao.SysAftersales.Columns().OrderId:    req.Id,
			dao.SysAftersales.Columns().Code:       utils_code.GetCode(ctx, consts.AS),
			dao.SysAftersales.Columns().Amount:     req.Amount,
			dao.SysAftersales.Columns().Type:       req.Type,
			dao.SysAftersales.Columns().ManageId:   ctx.Value("userId"),
			dao.SysAftersales.Columns().Reason:     req.Reason,
			dao.SysAftersales.Columns().Status:     consts.StatusSuccess,
			dao.SysAftersales.Columns().CreateTime: gtime.Now(),
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
		dao.SysCapital.Columns().Type:       consts.CapitalRefundOrder,
		dao.SysCapital.Columns().Mode:       consts.PayModePersonalTransfer,
		dao.SysCapital.Columns().UserId:     order.GMap().Get(dao.SysOrder.Columns().UserId),
	}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	orderUserId := gconv.Int64(order.GMap().Get(dao.SysOrder.Columns().UserId))
	_, err = tx.Model(dao.SysUserBill.Table()).
		Data(g.Map{
			dao.SysUserBill.Columns().UserId:     orderUserId,
			dao.SysUserBill.Columns().RelatedId:  req.Id,
			dao.SysUserBill.Columns().Code:       utils_code.GetCode(ctx, consts.BL),
			dao.SysUserBill.Columns().Type:       consts.BillTypeRefund,
			dao.SysUserBill.Columns().Amount:     req.Amount,
			dao.SysUserBill.Columns().Mode:       consts.Add,
			dao.SysUserBill.Columns().CreateTime: gtime.Now(),
		}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	return
}
