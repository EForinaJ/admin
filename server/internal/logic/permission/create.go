package permission

import (
	"context"

	"server/internal/dao"
	"server/internal/model/do"
	dto_permission "server/internal/type/permission/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Create implements service.IPermission.
func (s *sPermission) Create(ctx context.Context, req *dto_permission.Create) (err error) {
	var entity *do.SysPermission
	err = gconv.Struct(req, &entity)
	if err != nil {
		return utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}

	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()

	_, err = dao.SysPermission.Ctx(ctx).Insert(&entity)
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	return
}
