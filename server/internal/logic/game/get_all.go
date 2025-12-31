package game

import (
	"context"

	"server/internal/dao"
	"server/internal/model/entity"
	dao_game "server/internal/type/game/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetAll implements service.IGame.
func (s *sGame) GetAll(ctx context.Context) (res []*dao_game.List, err error) {
	var list []*entity.SysGame
	err = dao.SysGame.Ctx(ctx).Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_game.List, len(list))
	for i, v := range list {
		var entity *dao_game.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return nil, utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
		}
		res[i] = entity
	}

	return
}
