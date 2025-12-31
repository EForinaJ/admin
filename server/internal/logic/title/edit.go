package title

import (
	"context"
	"server/internal/dao"
	"server/internal/model/do"
	dto_title "server/internal/type/title/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Edit implements service.ITitle.
func (s *sTitle) Edit(ctx context.Context, req *dto_title.Edit) (err error) {

	var entity *do.SysTitle
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}
	entity.Name = req.Name
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysTitle.Ctx(ctx).
		Where(dao.SysTitle.Columns().Id, req.Id).
		Data(&entity).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR, response.CodeMsg(response.DB_SAVE_ERROR))
	}
	return
}
