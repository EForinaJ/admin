package role

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_role "server/internal/type/role/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetAll implements service.IRole.
func (s *sRole) GetAll(ctx context.Context) (res []*dao_role.List, err error) {
	m := dao.SysRole.Ctx(ctx).
		OrderDesc(dao.SysRole.Columns().CreateTime)

	var list []*entity.SysRole
	err = m.Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_role.List, len(list))
	for i, v := range list {
		var entity *dao_role.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return nil, utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
		}
		res[i] = entity
	}

	return
}
