package level

import (
	"context"

	"server/internal/dao"
	dto_level "server/internal/type/level/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckCreate implements service.ILevel.
func (s *sLevel) CheckCreate(ctx context.Context, req *dto_level.Create) (err error) {
	res, err := dao.SysLevel.Ctx(ctx).
		Where(dao.SysLevel.Columns().Name, req.Name).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if res {
		return utils_error.Err(response.FAILD, "该等级已存在")
	}

	return
}
