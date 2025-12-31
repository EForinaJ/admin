package menu

import (
	"context"
	"server/internal/dao"
	dto_menu "server/internal/type/menu/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckCreate implements service.IMenu.
func (s *sMenu) CheckCreate(ctx context.Context, req *dto_menu.Create) (err error) {
	exist, err := dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Path, req.Path).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if exist {
		return utils_error.Err(response.FAILD, "菜单路径已存在")
	}

	return
}
