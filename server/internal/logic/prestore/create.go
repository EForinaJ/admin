package prestore

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_prestore "server/internal/type/prestore/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Create implements service.IPrestore.
func (s *sPrestore) Create(ctx context.Context, req *dto_prestore.Create) (err error) {
	//  添加预存
	_, err = dao.SysPrestore.Ctx(ctx).Data(g.Map{
		dao.SysPrestore.Columns().CreateTime:  gtime.Now(),
		dao.SysPrestore.Columns().ManageId:    ctx.Value("userId"),
		dao.SysPrestore.Columns().BonusAmount: req.BonusAmount,
		dao.SysPrestore.Columns().Amount:      req.Amount,
		dao.SysPrestore.Columns().UserId:      req.UserId,
		dao.SysPrestore.Columns().Status:      consts.StatusApply,
	}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	return
}
