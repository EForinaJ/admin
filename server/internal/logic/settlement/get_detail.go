package settlement

import (
	"context"
	"server/internal/dao"
	"server/internal/service"
	dao_settlement "server/internal/type/settlement/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.ISettlement.
func (s *sSettlement) GetDetail(ctx context.Context, id int64) (res *dao_settlement.Detail, err error) {
	obj, err := dao.SysSettlement.Ctx(ctx).
		WherePri(id).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if obj.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}
	// 2. 使用结构体转换代替手动映射
	var detail *dao_settlement.Detail
	if err := gconv.Scan(obj.Map(), &detail); err != nil {
		return nil, utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}

	//  订单
	order, err := service.Order().GetDetail(ctx, gconv.Int64(obj.GMap().Get(dao.SysSettlement.Columns().OrderId)))
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Order = order

	//  威客
	witkey, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().Id, obj.GMap().Get(dao.SysSettlement.Columns().WitkeyId)).
		Fields(dao.SysWitkey.Columns().UserId, dao.SysWitkey.Columns().TitleId, dao.SysWitkey.Columns().GameId).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	user, err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id,
			witkey.GMap().Get(dao.SysWitkey.Columns().UserId)).
		Value(dao.SysUser.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Witkey = user.String()

	manage, err := dao.SysManage.Ctx(ctx).
		Where(dao.SysManage.Columns().Id,
			obj.GMap().Get(dao.SysSettlement.Columns().ManageId)).
		Value(dao.SysManage.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Manage = manage.String()

	return detail, nil
}
