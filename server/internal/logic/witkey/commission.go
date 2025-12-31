package witkey

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_witkey "server/internal/type/witkey/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

// ChangeCommission implements service.IWitkey.
func (s *sWitkey) ChangeCommission(ctx context.Context, req *dto_witkey.Commission) (err error) {
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
	commission, err := tx.Model(dao.SysWitkey.Table()).
		WherePri(req.Id).Value(dao.SysWitkey.Columns().Commission)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	amount := decimal.NewFromFloat(commission.Float64())
	entity := g.Map{
		dao.SysCommission.Columns().After:      commission,
		dao.SysCommission.Columns().Amount:     req.Amount,
		dao.SysCommission.Columns().Mode:       req.Mode,
		dao.SysCommission.Columns().WitkeyId:   req.Id,
		dao.SysCommission.Columns().CreateTime: gtime.Now(),
		dao.SysCommission.Columns().Type:       consts.WitkeyChangeCommissionTypeSystem,
	}

	switch req.Mode {
	case consts.Add:
		newCommission := amount.Add(decimal.NewFromFloat(req.Amount))
		_, err = tx.Model(dao.SysWitkey.Table()).WherePri(req.Id).Data(g.Map{
			dao.SysWitkey.Columns().Commission: newCommission,
		}).Update()
		if err != nil {
			return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}
		entity[dao.SysCommission.Columns().Before] = newCommission
		entity[dao.SysCommission.Columns().Remark] = "系统增加余额"
		entity[dao.SysCommission.Columns().Related] = "系统增加余额"
	case consts.Sub:
		newCommission := amount.Sub(decimal.NewFromFloat(req.Amount))
		_, err = tx.Model(dao.SysWitkey.Table()).WherePri(req.Id).Data(g.Map{
			dao.SysWitkey.Columns().Commission: newCommission,
		}).Update()
		if err != nil {
			return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}
		entity[dao.SysCommission.Columns().Related] = "系统减少余额"
		entity[dao.SysCommission.Columns().Remark] = "系统减少余额"
		entity[dao.SysCommission.Columns().Before] = newCommission
	}

	if req.Remark != "" {
		entity[dao.SysCommission.Columns().Remark] = req.Remark
	}

	_, err = tx.Model(dao.SysCommission.Table()).
		Data(entity).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	return
}
