package aftersales

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_aftersales "server/internal/type/aftersales/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckApply implements service.IAftersales.
func (s *sAftersales) CheckApply(ctx context.Context, req *dto_aftersales.Apply) (err error) {
	status, err := dao.SysAftersales.Ctx(ctx).Where(dao.SysAftersales.Columns().Id, req.Id).
		Value(dao.SysAftersales.Columns().Status)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if status.Int() != consts.StatusApply {
		return utils_error.Err(response.FAILD, "该售后工单已审核")
	}
	return
}
