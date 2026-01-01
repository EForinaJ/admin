package site

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_site "server/internal/type/site/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// GetUserOptions implements service.ISite.
func (s *sSite) GetUserOptions(ctx context.Context, phone string) (res []*dao_site.Options, err error) {
	m := dao.SysUser.Ctx(ctx).
		Page(1, 5).
		OrderDesc(dao.SysUser.Columns().CreateTime).
		Fields(dao.SysUser.Columns().Id, dao.SysUser.Columns().Name)
	m = m.Where(dao.SysUser.Columns().Phone, phone)

	var list []*entity.SysUser
	err = m.Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	res = make([]*dao_site.Options, len(list))
	for i, v := range list {
		res[i] = &dao_site.Options{
			Name:  v.Name,
			Id:    v.Id,
			Value: v.Id,
		}
	}
	return
}
