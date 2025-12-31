package role

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_role "server/internal/type/role/dao"
	dto_role "server/internal/type/role/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IRole.
func (s *sRole) GetList(ctx context.Context, req *dto_role.Query) (total int, res []*dao_role.List, err error) {
	m := dao.SysRole.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysRole.Columns().CreateTime)

	if req.Name != "" {
		m = m.Where(dao.SysRole.Columns().Name, req.Name)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var list []*entity.SysRole
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_role.List, len(list))
	for i, v := range list {
		var entity *dao_role.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		res[i] = entity
	}

	return
}
