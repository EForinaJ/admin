package settlement

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_settlement "server/internal/type/settlement/dto"
	utils_code "server/internal/utils/code"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// Apply implements service.ISettlement.
func (s *sSettlement) Apply(ctx context.Context, req *dto_settlement.Apply) (err error) {
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
	obj, err := tx.Model(dao.SysSettlement.Table()).
		Where(dao.SysSettlement.Columns().Id, req.Id).
		Fields(
			dao.SysSettlement.Columns().WitkeyId,
			dao.SysSettlement.Columns().Commission,
			dao.SysSettlement.Columns().ServiceCharge,
			dao.SysSettlement.Columns().OrderId,
		).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	witkeyCommission, err := tx.Model(dao.SysWitkey.Table()).
		Where(dao.SysWitkey.Columns().Id, obj.GMap().Get(dao.SysSettlement.Columns().WitkeyId)).
		Value(dao.SysWitkey.Columns().Commission)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	commission := decimal.NewFromFloat(gconv.Float64(obj.GMap().Get(dao.SysSettlement.Columns().Commission)))
	//  同意结算
	if req.Status == consts.StatusSuccess {

		_, err = tx.Model(dao.SysSettlement.Table()).
			Where(dao.SysSettlement.Columns().Id, req.Id).
			Data(g.Map{
				dao.SysSettlement.Columns().Status:         consts.StatusSuccess,
				dao.SysSettlement.Columns().SettlementTime: gtime.Now(),
			}).Update()
		if err != nil {
			return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}

		newCommission := decimal.NewFromFloat(witkeyCommission.Float64()).Add(commission)
		_, err = tx.Model(dao.SysWitkey.Table()).
			Where(dao.SysWitkey.Columns().Id, obj.GMap().Get(dao.SysSettlement.Columns().WitkeyId)).
			Data(g.Map{
				dao.SysWitkey.Columns().Commission: newCommission,
			}).Update()
		if err != nil {
			return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}

		_, err = tx.Model(dao.SysCommission.Table()).Data(g.Map{
			dao.SysCommission.Columns().WitkeyId:   obj.GMap().Get(dao.SysSettlement.Columns().WitkeyId),
			dao.SysCommission.Columns().Before:     witkeyCommission,
			dao.SysCommission.Columns().After:      newCommission,
			dao.SysCommission.Columns().Amount:     commission,
			dao.SysCommission.Columns().Mode:       consts.Add,
			dao.SysCommission.Columns().CreateTime: gtime.Now(),
			dao.SysCommission.Columns().Type:       consts.WitkeyChangeCommissionTypeSettlement,
			dao.SysCommission.Columns().Remark:     "报单结算变更佣金",
			dao.SysCommission.Columns().Related:    obj.GMap().Get(dao.SysSettlement.Columns().Code),
		}).Insert()
		if err != nil {
			return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
		}
		userId, err := tx.Model(dao.SysWitkey.Table()).
			Where(dao.SysWitkey.Columns().Id, obj.GMap().Get(dao.SysSettlement.Columns().WitkeyId)).
			Value(dao.SysWitkey.Columns().UserId)
		if err != nil {
			return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		_, err = tx.Model(dao.SysUserBill.Table()).
			Data(g.Map{
				dao.SysUserBill.Columns().UserId:     userId,
				dao.SysUserBill.Columns().RelatedId:  req.Id,
				dao.SysUserBill.Columns().Code:       utils_code.GetCode(ctx, consts.BL),
				dao.SysUserBill.Columns().Type:       consts.BillTypeSettlementCommission,
				dao.SysUserBill.Columns().Amount:     commission,
				dao.SysUserBill.Columns().Mode:       consts.Add,
				dao.SysUserBill.Columns().CreateTime: gtime.Now(),
			}).Insert()
		if err != nil {
			return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
		}

		orderCode, err := tx.Model(dao.SysOrder.Table()).
			Where(dao.SysOrder.Columns().Id, obj.GMap().Get(dao.SysSettlement.Columns().OrderId)).
			Value(dao.SysOrder.Columns().Code)
		if err != nil {
			return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		_, err = tx.Model(dao.SysProfit.Table()).
			Data(g.Map{
				dao.SysProfit.Columns().Type:       userId,
				dao.SysProfit.Columns().Related:    orderCode.String(),
				dao.SysProfit.Columns().Type:       consts.ProfitTypeOrder,
				dao.SysProfit.Columns().Amount:     obj.GMap().Get(dao.SysSettlement.Columns().ServiceCharge),
				dao.SysProfit.Columns().CreateTime: gtime.Now(),
			}).Insert()
		if err != nil {
			return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
		}
	}

	// 拒绝结算
	if req.Status == consts.StatusFail {
		_, err = tx.Model(dao.SysSettlement.Table()).
			Where(dao.SysSettlement.Columns().Id, req.Id).
			Data(g.Map{
				dao.SysSettlement.Columns().Reason: req.Reason,
				dao.SysSettlement.Columns().Status: consts.StatusFail,
			}).Update()
		if err != nil {
			return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}
	}

	return
}
