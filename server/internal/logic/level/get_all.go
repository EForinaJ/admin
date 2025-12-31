package level

import (
	"context"

	"server/internal/dao"
	"server/internal/model/entity"
	dao_level "server/internal/type/level/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetAll implements service.ILevel.
func (s *sLevel) GetAll(ctx context.Context) (res []*dao_level.List, err error) {
	var list []*entity.SysLevel
	err = dao.SysLevel.Ctx(ctx).Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_level.List, len(list))
	for i, v := range list {
		var entity *dao_level.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return nil, utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
		}
		res[i] = entity
	}

	return
}
