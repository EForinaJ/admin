package prestore

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_prestore "server/internal/type/prestore/dao"
	dto_prestore "server/internal/type/prestore/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IPrestore.
func (s *sPrestore) GetList(ctx context.Context, req *dto_prestore.Query) (total int, res []*dao_prestore.List, err error) {
	m := dao.SysPrestore.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysPrestore.Columns().CreateTime)
	if req.Name != "" {
		userId, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Name, req.Name).
			Value(dao.SysUser.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		m = m.Where(dao.SysPrestore.Columns().UserId, userId)
	}
	if req.Status != 0 {
		m = m.Where(dao.SysPrestore.Columns().Status, req.Status)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var list []*entity.SysPrestore
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_prestore.List, len(list))
	for i, v := range list {
		var entity *dao_prestore.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		user, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Id, v.UserId).
			Value(dao.SysUser.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.User = user.String()

		res[i] = entity
	}

	return
}
