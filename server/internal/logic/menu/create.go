package menu

import (
	"context"

	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/do"
	dto_menu "server/internal/type/menu/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// Create implements service.IMenu.
func (s *sMenu) Create(ctx context.Context, req *dto_menu.Create) (err error) {
	var entity *do.SysMenu
	err = gconv.Struct(req, &entity)
	if err != nil {
		return utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}
	entity.Component = req.Component

	// 布尔值映射优化
	boolMapping := map[bool]int{
		true:  consts.Yes,
		false: consts.Not,
	}
	entity.IsEnable = boolMapping[req.IsEnable]
	entity.IsMenu = boolMapping[req.IsMenu]
	entity.KeepAlive = boolMapping[req.KeepAlive]
	entity.IsHide = boolMapping[req.IsHide]
	entity.IsHideTab = boolMapping[req.IsHideTab]
	entity.IsIframe = boolMapping[req.IsIframe]
	entity.ShowBadge = boolMapping[req.ShowBadge]
	entity.ShowTextBadge = boolMapping[req.ShowTextBadge]
	entity.FixedTab = boolMapping[req.FixedTab]
	entity.IsFullPage = boolMapping[req.IsFullPage]
	_, err = dao.SysMenu.Ctx(ctx).Insert(&entity)
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	return
}
