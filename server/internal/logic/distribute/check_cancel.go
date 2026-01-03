package distribute

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_distribute "server/internal/type/distribute/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckCancel implements service.IDistribute.
func (s *sDistribute) CheckCancel(ctx context.Context, req *dto_distribute.Cancel) (err error) {
	exist, err := dao.SysDistribute.Ctx(ctx).
		Where(dao.SysDistribute.Columns().Id, req.Id).
		Where(dao.SysDistribute.Columns().Status, consts.DistributeStatusCancel).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if exist {
		return utils_error.Err(response.FAILD, "该派单已经取消过，请勿重复操作")
	}
	return
}
