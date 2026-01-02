package prestore

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_prestore "server/internal/type/prestore/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckApply implements service.IPrestore.
func (s *sPrestore) CheckApply(ctx context.Context, req *dto_prestore.Apply) (err error) {
	status, err := dao.SysPrestore.Ctx(ctx).
		WherePri(req.Id).
		Value(dao.SysPrestore.Columns().Status)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if status.Int() != consts.StatusApply {
		return utils_error.Err(response.FAILD, "预存申请已经审核")
	}

	return
}
