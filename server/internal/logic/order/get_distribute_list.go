package order

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_order "server/internal/type/order/dao"
	dto_order "server/internal/type/order/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDistributeList implements service.IOrder.
func (s *sOrder) GetDistributeList(ctx context.Context, req *dto_order.DistributeQuery) (total int, res []*dao_order.DistributeList, err error) {
	m := dao.SysDistribute.Ctx(ctx).
		Page(req.Page, req.Limit).
		Where(dao.SysDistribute.Columns().OrderId, req.Id).
		OrderDesc(dao.SysDistribute.Columns().CreateTime)
	if req.Name != "" {
		userId, err := dao.SysUser.Ctx(ctx).
			WhereLike(dao.SysUser.Columns().Name, "%"+req.Name+"%").Value(dao.SysUser.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		witkeyId, err := dao.SysWitkey.Ctx(ctx).
			Where(dao.SysWitkey.Columns().UserId, userId).Value(dao.SysWitkey.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		m = m.Where(dao.SysDistribute.Columns().WitkeyId, witkeyId)
	}

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var list []*entity.SysDistribute
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_order.DistributeList, len(list))
	for i, v := range list {
		var entity *dao_order.DistributeList
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		witkey, err := dao.SysWitkey.Ctx(ctx).Where(dao.SysWitkey.Columns().Id, v.WitkeyId).
			Fields(dao.SysWitkey.Columns().UserId, dao.SysWitkey.Columns().TitleId, dao.SysWitkey.Columns().GameId).One()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		//  下单用户
		user, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Id, witkey.GMap().Get(dao.SysWitkey.Columns().UserId)).
			Value(dao.SysUser.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Witkey = user.String()

		game, err := dao.SysGame.Ctx(ctx).
			Where(dao.SysGame.Columns().Id,
				witkey.GMap().Get(dao.SysWitkey.Columns().GameId)).
			Value(dao.SysGame.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Game = game.String()

		title, err := dao.SysTitle.Ctx(ctx).
			Where(dao.SysTitle.Columns().Id,
				witkey.GMap().Get(dao.SysWitkey.Columns().TitleId)).
			Value(dao.SysTitle.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Title = title.String()

		res[i] = entity
	}
	return
}
