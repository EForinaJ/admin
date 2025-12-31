package witkey

import (
	"context"
	"server/internal/dao"
	dto_witkey "server/internal/type/witkey/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/shopspring/decimal"
)

// CheckCommission implements service.IWitkey.
func (s *sWitkey) CheckCommission(ctx context.Context, req *dto_witkey.Commission) (err error) {
	commission, err := dao.SysWitkey.Ctx(ctx).
		WherePri(req.Id).
		Value(dao.SysWitkey.Columns().Commission)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if !decimal.NewFromFloat(req.Amount).LessThanOrEqual(decimal.NewFromFloat(commission.Float64())) {
		return utils_error.Err(response.FAILD, "佣金不足，减少金额超出佣金")
	}
	return
}
