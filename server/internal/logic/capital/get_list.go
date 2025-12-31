package capital

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_capital "server/internal/type/capital/dao"
	dto_capital "server/internal/type/capital/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.ICapital.
func (s *sCapital) GetList(ctx context.Context, req *dto_capital.Query) (total int, res []*dao_capital.List, err error) {
	m := dao.SysCapital.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysCapital.Columns().CreateTime)
	if req.Code != "" {
		m = m.Where(dao.SysCapital.Columns().Code, req.Code)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var list []*entity.SysCapital
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_capital.List, len(list))
	for i, v := range list {
		var entity *dao_capital.List
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
