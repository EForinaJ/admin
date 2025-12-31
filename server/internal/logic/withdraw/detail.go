package withdraw

import (
	"context"
	"server/internal/dao"
	dao_withdraw "server/internal/type/withdraw/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IWithdraw.
func (s *sWithdraw) GetDetail(ctx context.Context, id int64) (res *dao_withdraw.Detail, err error) {
	info, err := dao.SysWithdraw.Ctx(ctx).
		Where(dao.SysWithdraw.Columns().Id, id).One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if info.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}

	var detail *dao_withdraw.Detail
	err = gconv.Scan(info, &detail)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	userId, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().Id, info.GMap().Get(dao.SysWithdraw.Columns().WitkeyId)).
		Value(dao.SysWitkey.Columns().UserId)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	user, err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id, userId).
		Value(dao.SysUser.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Witkey = user.String()

	//
	manage, err := dao.SysManage.Ctx(ctx).
		Where(dao.SysManage.Columns().Id, info.GMap().Get(dao.SysWithdraw.Columns().ManageId)).
		Value(dao.SysManage.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Manage = gconv.String(manage)

	return detail, nil
}
