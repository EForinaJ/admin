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
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// Create implements service.IUser.
func (s *sUser) Create(ctx context.Context, req *dto_user.Create) (err error) {

	var entity *do.SysUser
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	if req.Password == "" {
		req.Password = "123456"
	}
	newSalt := grand.S(6)
	newToken := consts.SYSTEMNAME + req.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)
	entity.Salt = newSalt
	entity.Password = newToken
	if req.Name == "" {
		req.Name = "新用户" + grand.S(6)
	}
	// 转换为 gtime.Time
	if req.Birthday == 0 {
		req.Birthday = 631123200000
	}
	birthday := gtime.NewFromTimeStamp(req.Birthday / 1000).UTC()
	entity.Birthday = birthday
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()

	//  获取配置
	userInfo, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.UserSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	userJosn := gjson.New(userInfo)
	//  获取默认头像和封面
	if req.Avatar == "" {
		entity.Avatar = userJosn.Get("avatar")
	}

	entity.Cover = userJosn.Get("cover")
	entity.Status = req.Status

	entity.LevelId = userJosn.Get("levelId").Int64()
	exp, err := dao.SysLevel.Ctx(ctx).
		Where(dao.SysLevel.Columns().Id, userJosn.Get("levelId").Int64()).
		Value(dao.SysLevel.Columns().Experience)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	entity.Experience = exp.Int()

	_, err = dao.SysUser.Ctx(ctx).Data(&entity).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	// redis 新增

	return
}
