package level

import (
	"context"

	"server/internal/dao"
	dao_level "server/internal/type/level/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// GetEdit implements service.ILevel.
func (s *sLevel) GetEdit(ctx context.Context, id int64) (res *dao_level.Edit, err error) {
	err = dao.SysLevel.Ctx(ctx).Where(dao.SysLevel.Columns().Id, id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	return
}
