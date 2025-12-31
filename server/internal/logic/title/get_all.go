package title

import (
	"context"

	"server/internal/dao"
	dao_title "server/internal/type/title/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// GetAll implements service.ITitle.
func (s *sTitle) GetAll(ctx context.Context, gameId int64) (res []*dao_title.OptionsList, err error) {
	err = dao.SysTitle.Ctx(ctx).Where(dao.SysTitle.Columns().GameId, gameId).
		Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	return
}
