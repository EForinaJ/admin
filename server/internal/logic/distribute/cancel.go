package distribute

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_distribute "server/internal/type/distribute/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Cancel implements service.IOrder.
func (s *sDistribute) Cancel(ctx context.Context, req *dto_distribute.Cancel) (err error) {
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
	_, err = tx.Model(dao.SysDistribute.Table()).
		Where(dao.SysDistribute.Columns().Id, req.Id).
		Data(g.Map{
			dao.SysDistribute.Columns().Status: consts.DistributeStatusCancel,
			dao.SysDistribute.Columns().Reason: req.Reason,
		}).
		Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}
	orderId, err := tx.Model(dao.SysDistribute.Table()).
		Where(dao.SysDistribute.Columns().Id, req.Id).Value(dao.SysDistribute.Columns().OrderId)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	//  添加订单日志
	_, err = tx.Model(dao.SysOrderLog.Table()).Data(g.Map{
		dao.SysOrderLog.Columns().CreateTime: gtime.Now(),
		dao.SysOrderLog.Columns().Content:    "取消派单服务",
		dao.SysOrderLog.Columns().ManageId:   ctx.Value("userId"),
		dao.SysOrderLog.Columns().OrderId:    orderId,
		dao.SysOrderLog.Columns().Type:       consts.OrderLogTypeDistributeCancel,
	}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	return
}
