package manage

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/do"
	dto_manage "server/internal/type/manage/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// Create implements service.IManage.
func (s *sManage) Create(ctx context.Context, req *dto_manage.Create) (err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return utils_error.Err(response.DB_TX_ERROR, response.CodeMsg(response.DB_TX_ERROR))
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var entity *do.SysManage
	err = gconv.Struct(req, &entity)
	if err != nil {
		return utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}
	if req.Password == "" {
		req.Password = grand.S(6)
	}
	newSalt := grand.S(6)
	newToken := consts.SYSTEMNAME + req.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)

	entity.Salt = newSalt
	entity.Password = newToken

	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	rs, err := tx.Model(dao.SysManage.Table()).
		Data(&entity).
		Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	rid, err := rs.LastInsertId()

	if err != nil || rid <= 0 {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	if len(req.Roles) > 0 {

		_, err = tx.Model(dao.SysManageRole.Table()).
			Where(dao.SysManageRole.Columns().ManageId, rid).Delete()
		if err != nil {
			return utils_error.Err(response.DELETE_FAILED, response.CodeMsg(response.DELETE_FAILED))
		}

		manageRoleEntites := make([]*do.SysManageRole, len(req.Roles))
		for i, v := range req.Roles {
			manageRoleEntites[i] = &do.SysManageRole{
				RoleId:   v,
				ManageId: rid,
			}
		}
		_, err = tx.Model(dao.SysManageRole.Table()).
			Data(&manageRoleEntites).
			Insert()
		if err != nil {
			return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
		}

		//  更新casbin 权限内容
		_, err = tx.Model(dao.SysCasbin.Table()).
			Where(dao.SysCasbin.Columns().PType, "g").
			Where(dao.SysCasbin.Columns().V0, rid).Delete()
		if err != nil {
			return utils_error.Err(response.DELETE_FAILED, response.CodeMsg(response.DELETE_FAILED))
		}

		roles, err := tx.Model(dao.SysRole.Table()).
			WhereIn(dao.SysRole.Columns().Id, req.Roles).
			Fields(dao.SysRole.Columns().Code).
			Array()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR, response.CodeMsg(response.DB_SAVE_ERROR))
		}

		casbinEntites := make([]g.Map, len(req.Roles))
		for i, v := range roles {
			casbinEntites[i] = g.Map{
				dao.SysCasbin.Columns().PType: "g",
				dao.SysCasbin.Columns().V0:    rid,
				dao.SysCasbin.Columns().V1:    v.String(),
			}
		}

		_, err = tx.Model(dao.SysCasbin.Table()).
			Data(casbinEntites).
			Insert()
		if err != nil {
			return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
		}
	}

	return
}
