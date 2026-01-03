package distribute

import (
	"context"
	"server/internal/dao"
	"server/internal/service"
	dao_distribute "server/internal/type/distribute/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IDistribute.
func (s *sDistribute) GetDetail(ctx context.Context, id int64) (res *dao_distribute.Detail, err error) {
	obj, err := dao.SysDistribute.Ctx(ctx).
		WherePri(id).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if obj.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}

	// 2. 使用结构体转换代替手动映射
	var detail *dao_distribute.Detail
	if err := gconv.Scan(obj.Map(), &detail); err != nil {
		return nil, utils_error.Err(response.FAILD, response.CodeMsg(response.FAILD))
	}

	order, err := service.Order().GetDetail(ctx, gconv.Int64(obj.GMap().Get(dao.SysDistribute.Columns().OrderId)))
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Order = order

	witkey, err := dao.SysWitkey.Ctx(ctx).Where(dao.SysWitkey.Columns().Id, obj.GMap().Get(dao.SysDistribute.Columns().WitkeyId)).
		Fields(dao.SysWitkey.Columns().Name, dao.SysWitkey.Columns().TitleId, dao.SysWitkey.Columns().GameId).One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Witkey = gconv.String(witkey.GMap().Get(dao.SysWitkey.Columns().Name))

	game, err := dao.SysGame.Ctx(ctx).
		Where(dao.SysGame.Columns().Id,
			witkey.GMap().Get(dao.SysWitkey.Columns().GameId)).
		Value(dao.SysGame.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Game = game.String()

	title, err := dao.SysTitle.Ctx(ctx).
		Where(dao.SysTitle.Columns().Id,
			witkey.GMap().Get(dao.SysWitkey.Columns().TitleId)).
		Value(dao.SysTitle.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Title = title.String()

	manage, err := dao.SysManage.Ctx(ctx).
		Where(dao.SysManage.Columns().Id,
			obj.GMap().Get(dao.SysDistribute.Columns().ManageId)).
		Value(dao.SysManage.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Manage = manage.String()

	return detail, nil
}
