package title

import (
	"context"

	"server/internal/dao"
	dto_title "server/internal/type/title/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckEdit implements service.ITitle.
func (s *sTitle) CheckEdit(ctx context.Context, req *dto_title.Edit) (err error) {
	res, err := dao.SysTitle.Ctx(ctx).
		Where(dao.SysTitle.Columns().GameId, req.GameId).
		Where(dao.SysTitle.Columns().Name, req.Name).
		WhereNotIn(dao.SysTitle.Columns().Id, req.Id).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if res {
		return utils_error.Err(response.FAILD, "头衔已存在")
	}
	return
}
