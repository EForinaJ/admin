package user

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_user "server/internal/type/user/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

// ChangeBalance implements service.IUser.
func (s *sUser) ChangeBalance(ctx context.Context, req *dto_user.ChangeBalance) (err error) {

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
	balance, err := tx.Model(dao.SysUser.Table()).
		WherePri(req.Id).Value(dao.SysUser.Columns().Balance)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	amount := decimal.NewFromFloat(balance.Float64())
	entity := g.Map{
		dao.SysBalance.Columns().After:      balance,
		dao.SysBalance.Columns().Amount:     req.Amount,
		dao.SysBalance.Columns().Mode:       req.Mode,
		dao.SysBalance.Columns().UserId:     req.Id,
		dao.SysBalance.Columns().CreateTime: gtime.Now(),
		dao.SysBalance.Columns().Type:       consts.UserChangeBalanceTypeSystem,
	}

	switch req.Mode {
	case consts.Add:
		newBalance := amount.Add(decimal.NewFromFloat(req.Amount))
		_, err = tx.Model(dao.SysUser.Table()).WherePri(req.Id).Data(g.Map{
			dao.SysUser.Columns().Balance: newBalance,
		}).Update()
		if err != nil {
			return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}
		entity[dao.SysBalance.Columns().Before] = newBalance
		entity[dao.SysBalance.Columns().Remark] = "系统增加余额"
		entity[dao.SysBalance.Columns().Related] = "系统增加余额"
	case consts.Sub:
		newBalance := amount.Sub(decimal.NewFromFloat(req.Amount))
		_, err = tx.Model(dao.SysUser.Table()).WherePri(req.Id).Data(g.Map{
			dao.SysUser.Columns().Balance: newBalance,
		}).Update()
		if err != nil {
			return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}
		entity[dao.SysBalance.Columns().Related] = "系统减少余额"
		entity[dao.SysBalance.Columns().Remark] = "系统减少余额"
		entity[dao.SysBalance.Columns().Before] = newBalance
	}

	if req.Remark != "" {
		entity[dao.SysBalance.Columns().Remark] = req.Remark
	}

	_, err = tx.Model(dao.SysBalance.Table()).
		Data(entity).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	return
}
