package order

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_order "server/internal/type/order/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckDistributeCancel implements service.IOrder.
func (s *sOrder) CheckDistributeCancel(ctx context.Context, req *dto_order.DistributeCancel) (err error) {
	exist, err := dao.SysDistribute.Ctx(ctx).
		Where(dao.SysDistribute.Columns().Id, req.Id).
		Where(dao.SysDistribute.Columns().IsCancel, consts.Yes).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if exist {
		return utils_error.Err(response.FAILD, "该派单已经取消过，请勿重复操作")
	}
	return
}
