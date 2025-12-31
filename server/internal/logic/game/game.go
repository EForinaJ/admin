package game

import (
	"context"

	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"
	dao_game "server/internal/type/game/dao"
	dto_game "server/internal/type/game/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

type sGame struct{}

// Delete implements service.IGame.
func (s *sGame) Delete(ctx context.Context, ids []int64) (err error) {

	_, err = dao.SysGame.Ctx(ctx).
		WhereIn(dao.SysGame.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED, response.CodeMsg(response.DELETE_FAILED))
	}
	return
}

// GetEdit implements service.IGame.
func (s *sGame) GetEdit(ctx context.Context, id int64) (res *dao_game.Edit, err error) {
	err = dao.SysGame.Ctx(ctx).Where(dao.SysGame.Columns().Id, id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	return
}

// GetList implements service.IGame.
func (s *sGame) GetList(ctx context.Context, req *dto_game.Query) (total int, res []*dao_game.List, err error) {
	m := dao.SysGame.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysGame.Columns().CreateTime)

	if req.Name != "" {
		m = m.Where(dao.SysGame.Columns().Name, req.Name)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var list []*entity.SysGame
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_game.List, len(list))
	for i, v := range list {
		var entity *dao_game.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		// 获取威客人数
		res[i] = entity
	}

	return
}

func init() {
	service.RegisterGame(&sGame{})
}
