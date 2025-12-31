package order

import (
	"context"

	"server/internal/dao"
	dto_order "server/internal/type/order/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/shopspring/decimal"
)

// AddDiscount implements service.IOrder.
func (s *sOrder) AddDiscount(ctx context.Context, req *dto_order.AddDiscount) (err error) {
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

	orderTotal, err := dao.SysOrder.Ctx(ctx).Where(dao.SysOrder.Columns().Id, req.Id).
		Value(dao.SysOrder.Columns().TotalAmount)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	discountAmout := decimal.NewFromFloat(req.Amount)
	actualAmount := decimal.NewFromFloat(orderTotal.Float64()).Sub(discountAmout)

	_, err = tx.Model(dao.SysOrder.Table()).WherePri(req.Id).Data(g.Map{
		dao.SysOrder.Columns().DiscountAmount: req.Amount,
		dao.SysOrder.Columns().ActualAmount:   actualAmount,
	}).Update()
	if err != nil {
		return utils_error.Err(response.UNIQUE_ERROR, response.CodeMsg(response.UNIQUE_ERROR))
	}

	return
}
