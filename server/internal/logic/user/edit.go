package user

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/do"
	dto_user "server/internal/type/user/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// Edit implements service.IUser.
func (s *sUser) Edit(ctx context.Context, req *dto_user.Edit) (err error) {

	var entity *do.SysUser
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}
	if req.Password != "" {
		newSalt := grand.S(6)
		newToken := consts.SYSTEMNAME + req.Password + newSalt
		newToken = gmd5.MustEncryptString(newToken)
		entity.Salt = newSalt
		entity.Password = newToken
	} else {
		password, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Id, req.Id).Value(dao.SysUser.Columns().Password)
		if err != nil {
			return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Password = password.String()
	}

	// 转换为 gtime.Time
	if req.Birthday == 0 {
		req.Birthday = 631123200000
	}
	birthday := gtime.NewFromTimeStamp(req.Birthday / 1000).UTC()
	entity.Birthday = birthday

	// userLevel, err := tx.Model(dao.SysUser.Table()).
	// 	Where(dao.SysUser.Columns().Id, req.Id).
	// 	Value(dao.SysUser.Columns().LevelId)
	// if err != nil {
	// 	return utils_error.Err(response.DB_SAVE_ERROR)
	// }
	// if userLevel.Int64() != req.LevelId {
	// 	exp, err := tx.Model(dao.SysLevel.Table()).
	// 		Where(dao.SysLevel.Columns().Id, req.LevelId).
	// 		Value(dao.SysLevel.Columns().Experience)
	// 	if err != nil {
	// 		return utils_error.Err(response.DB_SAVE_ERROR)
	// 	}
	// 	entity.Experience = exp.Int()
	// }

	entity.UpdateTime = gtime.Now()
	_, err = dao.SysUser.Ctx(ctx).
		WherePri(req.Id).
		Data(&entity).Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}

	return
}
