package permission

import (
	"context"

	"server/internal/dao"
	"server/internal/model/do"
	dto_permission "server/internal/type/permission/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Edit implements service.IPermission.
func (s *sPermission) Edit(ctx context.Context, req *dto_permission.Edit) (err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	pers, err := tx.Model(dao.SysPermission.Table()).
		Fields(dao.SysPermission.Columns().Permission).
		WhereIn(dao.SysPermission.Columns().Id, req.Id).Array()
	if err != nil {
		return err
	}

	_, err = tx.Model(dao.SysCasbin.Table()).
		WhereIn(dao.SysCasbin.Columns().V1, pers).
		Update(g.Map{
			dao.SysCasbin.Columns().V1: req.Permission,
		})
	if err != nil {
		return err
	}

	var entity *do.SysPermission
	err = gconv.Struct(req, &entity)
	if err != nil {
		return utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}
	entity.UpdateTime = gtime.Now()
	_, err = tx.Model(dao.SysPermission.Table()).
		WherePri(req.Id).
		Data(&entity).
		Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}

	return
}
