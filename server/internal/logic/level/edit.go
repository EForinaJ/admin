package level

import (
	"context"

	"server/internal/dao"
	"server/internal/model/do"
	dto_level "server/internal/type/level/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Edit implements service.ILevel.
func (s *sLevel) Edit(ctx context.Context, req *dto_level.Edit) (err error) {

	var entity *do.SysLevel
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysLevel.Ctx(ctx).
		Where(dao.SysLevel.Columns().Id, req.Id).
		Data(&entity).Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}
	return
}
