package distribute

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_distribute "server/internal/type/distribute/dao"
	dto_distribute "server/internal/type/distribute/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IDistribute.
func (s *sDistribute) GetList(ctx context.Context, req *dto_distribute.Query) (total int, res []*dao_distribute.List, err error) {
	m := dao.SysDistribute.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysDistribute.Columns().CreateTime)
	if req.Name != "" {
		witkeyId, err := dao.SysWitkey.Ctx(ctx).
			Where(dao.SysWitkey.Columns().Name, req.Name).Value(dao.SysWitkey.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		m = m.Where(dao.SysDistribute.Columns().WitkeyId, witkeyId)
	}
	if req.Code != "" {
		orderId, err := dao.SysOrder.Ctx(ctx).
			Where(dao.SysOrder.Columns().Code, req.Code).Value(dao.SysOrder.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		m = m.Where(dao.SysDistribute.Columns().OrderId, orderId)
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

	res = make([]*dao_distribute.List, len(list))
	for i, v := range list {
		var entity *dao_distribute.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		order, err := dao.SysOrder.Ctx(ctx).Where(dao.SysOrder.Columns().Id, v.OrderId).
			Value(dao.SysOrder.Columns().Code)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Order = order.String()

		witkey, err := dao.SysWitkey.Ctx(ctx).Where(dao.SysWitkey.Columns().Id, v.WitkeyId).
			Fields(dao.SysWitkey.Columns().Name, dao.SysWitkey.Columns().TitleId, dao.SysWitkey.Columns().GameId).One()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Witkey = gconv.String(witkey.GMap().Get(dao.SysWitkey.Columns().Name))

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

		manage, err := dao.SysManage.Ctx(ctx).
			Where(dao.SysManage.Columns().Id,
				v.ManageId).
			Value(dao.SysManage.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Manage = manage.String()

		res[i] = entity
	}
	return
}
