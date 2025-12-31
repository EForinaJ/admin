package system

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/do"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SaveConfig implements service.ISystem.
func (s *sSystem) SaveConfig(ctx context.Context, key string, name string, value interface{}) (err error) {
	entites := do.SysConfig{
		Key:   key,
		Name:  name,
		Value: value,
	}
	isExist, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, key).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	// 判断配置项是否存在
	if isExist {
		entites.UpdateTime = gtime.Now()
		_, err := g.Redis().Del(ctx, consts.BaseSetting)
		if err != nil {
			return utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
		}
		_, err = g.Redis().Del(ctx, "site")
		if err != nil {
			return utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
		}

		_, err = dao.SysConfig.Ctx(ctx).
			Where(dao.SysConfig.Columns().Key, key).
			Data(entites).Update()
		if err != nil {
			return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}
	} else {
		entites.CreateTime = gtime.Now()
		_, err = dao.SysConfig.Ctx(ctx).
			Data(entites).Insert()
		if err != nil {
			return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
		}
	}
	return
}
