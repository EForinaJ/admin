package title

import (
	"context"

	"server/internal/dao"
	dao_title "server/internal/type/title/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// GetEdit implements service.ITitle.
func (s *sTitle) GetEdit(ctx context.Context, id int64) (res *dao_title.Edit, err error) {
	err = dao.SysTitle.Ctx(ctx).Where(dao.SysTitle.Columns().Id, id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	return
}
