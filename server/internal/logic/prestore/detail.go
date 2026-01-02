package prestore

import (
	"context"
	"server/internal/dao"
	dao_prestore "server/internal/type/prestore/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IPrestore.
func (s *sPrestore) GetDetail(ctx context.Context, id int64) (res *dao_prestore.Detail, err error) {
	info, err := dao.SysPrestore.Ctx(ctx).
		Where(dao.SysPrestore.Columns().Id, id).One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if info.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}

	var detail *dao_prestore.Detail
	err = gconv.Scan(info, &detail)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	user, err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id, info.GMap().Get(dao.SysPrestore.Columns().UserId)).
		Value(dao.SysUser.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.User = user.String()

	//
	manage, err := dao.SysManage.Ctx(ctx).
		Where(dao.SysManage.Columns().Id, info.GMap().Get(dao.SysPrestore.Columns().ManageId)).
		Value(dao.SysManage.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Manage = gconv.String(manage)

	return detail, nil
}
