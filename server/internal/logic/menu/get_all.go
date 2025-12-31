package menu

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_menu "server/internal/type/menu/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetAll implements service.IMenu.
func (s *sMenu) GetAll(ctx context.Context) (res []*dao_menu.List, err error) {
	m := dao.SysMenu.Ctx(ctx).
		OrderDesc(dao.SysMenu.Columns().Sort)

	var list []*entity.SysMenu
	err = m.Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	res = make([]*dao_menu.List, len(list))
	// 布尔值映射优化
	boolMapping := map[int]bool{
		consts.Yes: true,
		consts.Not: false,
	}
	for i, v := range list {
		var entity *dao_menu.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.IsEnable = boolMapping[v.IsEnable]
		entity.IsMenu = boolMapping[v.IsMenu]
		entity.KeepAlive = boolMapping[v.KeepAlive]
		entity.IsHide = boolMapping[v.IsHide]
		entity.IsHideTab = boolMapping[v.IsHideTab]
		entity.IsIframe = boolMapping[v.IsIframe]
		entity.ShowBadge = boolMapping[v.ShowBadge]
		entity.ShowTextBadge = boolMapping[v.ShowTextBadge]
		entity.FixedTab = boolMapping[v.FixedTab]
		entity.IsFullPage = boolMapping[v.IsFullPage]

		res[i] = entity
	}
	return
}
