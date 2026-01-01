package order

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_order "server/internal/type/order/dao"
	dto_order "server/internal/type/order/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetLogList implements service.IOrder.
func (s *sOrder) GetLogList(ctx context.Context, req *dto_order.LogQuery) (total int, res []*dao_order.LogList, err error) {
	m := dao.SysOrderLog.Ctx(ctx).
		Page(req.Page, req.Limit).
		Where(dao.SysOrderLog.Columns().OrderId, req.Id).
		OrderDesc(dao.SysOrderLog.Columns().CreateTime)

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var list []*entity.SysOrderLog
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_order.LogList, len(list))
	for i, v := range list {
		var entity *dao_order.LogList
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		//  下单用户
		manage, err := dao.SysManage.Ctx(ctx).
			Where(dao.SysManage.Columns().Id, v.ManageId).
			Value(dao.SysManage.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Manage = manage.String()

		res[i] = entity
	}
	return
}
