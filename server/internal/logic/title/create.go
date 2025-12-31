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

// Create implements service.ITitle.
func (s *sTitle) Create(ctx context.Context, req *dto_title.Create) (err error) {
	var entity *do.SysTitle
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}
	entity.Name = req.Name
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysTitle.Ctx(ctx).Data(&entity).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	return
}
