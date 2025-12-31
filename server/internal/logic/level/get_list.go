package level

import (
	"context"

	"server/internal/dao"
	"server/internal/model/entity"
	dao_level "server/internal/type/level/dao"
	dto_level "server/internal/type/level/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.ILevel.
func (s *sLevel) GetList(ctx context.Context, req *dto_level.Query) (total int, res []*dao_level.List, err error) {
	m := dao.SysLevel.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysLevel.Columns().CreateTime)

	if req.Name != "" {
		m = m.Where(dao.SysLevel.Columns().Name, req.Name)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var list []*entity.SysLevel
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_level.List, len(list))
	for i, v := range list {
		var entity *dao_level.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		res[i] = entity
	}

	return
}
