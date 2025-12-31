package user

import (
	"context"
	"server/internal/dao"
	dao_user "server/internal/type/user/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetEdit implements service.IUser.
func (s *sUser) GetEdit(ctx context.Context, id int64) (res *dao_user.Edit, err error) {
	info, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, id).One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var detail *dao_user.Edit
	if err := gconv.Scan(info.Map(), &detail); err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	detail.Birthday = gtime.New(info.Map()[dao.SysUser.Columns().Birthday]).TimestampMilli()
	return detail, nil
}
