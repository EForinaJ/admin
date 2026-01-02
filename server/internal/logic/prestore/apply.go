package prestore

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_prestore "server/internal/type/prestore/dto"
	utils_code "server/internal/utils/code"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// Apply implements service.IPrestore.
func (s *sPrestore) Apply(ctx context.Context, req *dto_prestore.Apply) (err error) {
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

	if req.Status == consts.StatusSuccess {
		obj, err := tx.Model(dao.SysPrestore.Table()).
			Where(dao.SysPrestore.Columns().Id, req.Id).One()
		if err != nil {
			return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		amount := decimal.NewFromFloat(gconv.Float64(obj.GMap().Get(dao.SysPrestore.Columns().Amount))).
			Add(decimal.NewFromFloat(gconv.Float64(obj.GMap().Get(dao.SysPrestore.Columns().BonusAmount))))
		g.Dump(amount)
		code := utils_code.GetCode(ctx, consts.CZ)
		rs, err := tx.Model(dao.SysRecharge.Table()).
			Data(g.Map{
				dao.SysRecharge.Columns().Code:       code,
				dao.SysRecharge.Columns().Amount:     amount,
				dao.SysRecharge.Columns().PayMode:    consts.PayModePersonalTransfer,
				dao.SysRecharge.Columns().UserId:     obj.GMap().Get(dao.SysPrestore.Columns().UserId),
				dao.SysRecharge.Columns().Status:     consts.PayStatusSuccess,
				dao.SysRecharge.Columns().CreateTime: gtime.Now(),
			}).Insert()
		if err != nil {
			return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
		}

		rid, err := rs.LastInsertId()
		if err != nil {
			return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
		}

		userBalance, err := tx.Model(dao.SysUser.Table()).
			Where(dao.SysUser.Columns().Id, obj.GMap().Get(dao.SysPrestore.Columns().UserId)).Value(dao.SysUser.Columns().Balance)
		if err != nil {
			return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		balance := decimal.NewFromFloat(userBalance.Float64())
		newBalance := balance.Add(amount)
		_, err = tx.Model(dao.SysUser.Table()).
			Where(dao.SysUser.Columns().Id, obj.GMap().Get(dao.SysPrestore.Columns().UserId)).
			Data(g.Map{
				dao.SysUser.Columns().Balance: newBalance,
			}).Update()
		if err != nil {
			return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}

		billEntity := g.Map{
			dao.SysUserBill.Columns().UserId:     obj.GMap().Get(dao.SysPrestore.Columns().UserId),
			dao.SysUserBill.Columns().RelatedId:  rid,
			dao.SysUserBill.Columns().Code:       utils_code.GetCode(ctx, consts.BL),
			dao.SysUserBill.Columns().Type:       consts.BillTypeRecharge,
			dao.SysUserBill.Columns().Amount:     amount,
			dao.SysUserBill.Columns().Mode:       consts.Add,
			dao.SysUserBill.Columns().CreateTime: gtime.Now(),
		}
		_, err = tx.Model(dao.SysUserBill.Table()).
			Data(billEntity).Insert()
		if err != nil {
			return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
		}

		//  添加日志
		_, err = tx.Model(dao.SysBalance.Table()).Data(g.Map{
			dao.SysBalance.Columns().After:      balance,
			dao.SysBalance.Columns().Amount:     amount,
			dao.SysBalance.Columns().Before:     newBalance,
			dao.SysBalance.Columns().Mode:       consts.Add,
			dao.SysBalance.Columns().UserId:     obj.GMap().Get(dao.SysPrestore.Columns().UserId),
			dao.SysBalance.Columns().CreateTime: gtime.Now(),
			dao.SysBalance.Columns().Type:       consts.UserChangeBalanceTypeRecharge,
			dao.SysBalance.Columns().Remark:     "用户充值余额",
			dao.SysBalance.Columns().Related:    code,
		}).Insert()
		if err != nil {
			return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
		}
		//  添加支付日志
		_, err = tx.Model(dao.SysCapital.Table()).Data(g.Map{
			dao.SysCapital.Columns().CreateTime: gtime.Now(),
			dao.SysCapital.Columns().Code:       utils_code.GetCode(ctx, consts.PM),
			dao.SysCapital.Columns().Related:    code,
			dao.SysCapital.Columns().Amount:     amount,
			dao.SysCapital.Columns().Type:       consts.CapitalPaymentRecharge,
			dao.SysCapital.Columns().Mode:       consts.PayModePersonalTransfer,
			dao.SysCapital.Columns().UserId:     obj.GMap().Get(dao.SysPrestore.Columns().UserId),
		}).Insert()
		if err != nil {
			return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
		}

		_, err = tx.Model(dao.SysPrestore.Table()).WherePri(req.Id).Data(g.Map{
			dao.SysPrestore.Columns().Status: consts.StatusSuccess,
		}).Update()
		if err != nil {
			return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}

	}
	// 拒绝结算
	if req.Status == consts.StatusFail {
		_, err = tx.Model(dao.SysPrestore.Table()).
			Where(dao.SysPrestore.Columns().Id, req.Id).Data(g.Map{
			dao.SysPrestore.Columns().Status: consts.StatusFail,
			dao.SysPrestore.Columns().Reason: req.Reason,
		}).Update()
		if err != nil {
			return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}
	}

	return
}
