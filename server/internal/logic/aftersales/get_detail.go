package aftersales

import (
	"context"
	"server/internal/dao"
	"server/internal/service"
	dao_aftersales "server/internal/type/aftersales/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IAftersales.
func (s *sAftersales) GetDetail(ctx context.Context, id int64) (res *dao_aftersales.Detail, err error) {
	obj, err := dao.SysAftersales.Ctx(ctx).
		WherePri(id).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if obj.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}

	// 2. 使用结构体转换代替手动映射
	var detail *dao_aftersales.Detail
	if err := gconv.Scan(obj.Map(), &detail); err != nil {
		return nil, utils_error.Err(response.FAILD, response.CodeMsg(response.FAILD))
	}

	order, err := service.Order().GetDetail(ctx, gconv.Int64(obj.GMap().Get(dao.SysAftersales.Columns().OrderId)))
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.FAILD))
	}
	detail.Order = order

	manage, err := dao.SysManage.Ctx(ctx).
		Where(dao.SysManage.Columns().Id,
			obj.GMap().Get(dao.SysAftersales.Columns().ManageId)).
		Value(dao.SysManage.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Manage = manage.String()

	return detail, nil
}
