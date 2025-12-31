package game

import (
	"context"

	"server/internal/dao"
	dto_game "server/internal/type/game/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckCreate implements service.IGame.
func (s *sGame) CheckCreate(ctx context.Context, req *dto_game.Create) (err error) {
	res, err := dao.SysGame.Ctx(ctx).
		Where(dao.SysGame.Columns().Name, req.Name).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if res {
		return utils_error.Err(response.FAILD, "游戏名称已存在")
	}

	return
}
