package witkey

import (
	"context"
	"server/internal/dao"
	dto_witkey "server/internal/type/witkey/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
)

// ChangeTitle implements service.IWitkey.
func (s *sWitkey) ChangeTitle(ctx context.Context, req *dto_witkey.ChangeTitle) (err error) {
	_, err = dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().Id, req.Id).
		Data(g.Map{
			dao.SysWitkey.Columns().GameId:  req.GameId,
			dao.SysWitkey.Columns().TitleId: req.TitleId,
		}).Update()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	return
}
