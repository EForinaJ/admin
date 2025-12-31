package system

import (
	"context"
	"server/internal/dao"
	"server/internal/service"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

type sSystem struct{}

// GetBase implements service.ISystem.
func (s *sSystem) GetOne(ctx context.Context, key string) (res interface{}, err error) {
	res, err = dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, key).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	return
}

func init() {
	service.RegisterSystem(&sSystem{})
}
