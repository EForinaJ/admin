package game

import (
	"context"

	"server/internal/dao"
	"server/internal/model/do"
	dto_game "server/internal/type/game/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Create implements service.IGame.
func (s *sGame) Create(ctx context.Context, req *dto_game.Create) (err error) {

	var entity *do.SysGame
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysGame.Ctx(ctx).Data(&entity).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	return
}
