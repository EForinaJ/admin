package witkey

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_witkey "server/internal/type/witkey/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckChangeTitle implements service.IWitkey.
func (s *sWitkey) CheckChangeTitle(ctx context.Context, req *dto_witkey.ChangeTitle) (err error) {
	exist, err := dao.SysSettlement.Ctx(ctx).Where(dao.SysSettlement.Columns().WitkeyId, req.Id).
		Where(dao.SysSettlement.Columns().Status, consts.StatusApply).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if exist {
		return utils_error.Err(response.FAILD, "该威客还有报单未结算，无法更换头衔")
	}

	return
}
