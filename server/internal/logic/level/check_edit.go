package level

import (
	"context"
	"server/internal/dao"
	dto_level "server/internal/type/level/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckEdit implements service.ILevel.
func (s *sLevel) CheckEdit(ctx context.Context, req *dto_level.Edit) (err error) {
	res, err := dao.SysLevel.Ctx(ctx).
		Where(dao.SysLevel.Columns().Name, req.Name).
		WhereNotIn(dao.SysLevel.Columns().Id, req.Id).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if res {
		return utils_error.Err(response.FAILD, "该等级已存在")
	}

	return
}
