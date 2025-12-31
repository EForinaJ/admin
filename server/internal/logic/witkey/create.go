package witkey

import (
	"context"
	"server/internal/dao"
	dto_witkey "server/internal/type/witkey/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Create implements service.IWitkey.
func (s *sWitkey) Create(ctx context.Context, req *dto_witkey.Create) (err error) {
	_, err = dao.SysWitkey.Ctx(ctx).Data(g.Map{
		dao.SysWitkey.Columns().UserId:     req.UserId,
		dao.SysWitkey.Columns().GameId:     req.GameId,
		dao.SysWitkey.Columns().TitleId:    req.TitleId,
		dao.SysWitkey.Columns().Rate:       req.Rate,
		dao.SysWitkey.Columns().CreateTime: gtime.Now(),
		dao.SysWitkey.Columns().UpdateTime: gtime.Now(),
		dao.SysWitkey.Columns().Album:      req.Album,
	}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	return
}
