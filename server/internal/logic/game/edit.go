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

// Edit implements service.IGame.
func (s *sGame) Edit(ctx context.Context, req *dto_game.Edit) (err error) {

	var entity *do.SysGame
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysGame.Ctx(ctx).
		Where(dao.SysGame.Columns().Id, req.Id).
		Data(&entity).Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}

	return
}
