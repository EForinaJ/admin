package site

import (
	"context"

	"server/internal/dao"
	"server/internal/model/entity"
	dao_site "server/internal/type/site/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// GetCategoryOptions implements service.ISite.
func (s *sSite) GetCategoryOptions(ctx context.Context, id int64) (res []*dao_site.Options, err error) {
	m := dao.SysCategory.Ctx(ctx).
		OrderDesc(dao.SysCategory.Columns().CreateTime).
		Fields(dao.SysCategory.Columns().Id, dao.SysCategory.Columns().Name)
	m = m.Where(dao.SysCategory.Columns().GameId, id)

	var list []*entity.SysCategory
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
