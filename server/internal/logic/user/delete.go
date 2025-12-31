package user

import (
	"context"
	"server/internal/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
)

// Delete implements service.IUser.
func (s *sUser) Delete(ctx context.Context, ids []int64) (err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return utils_error.Err(response.DB_TX_ERROR, response.CodeMsg(response.DB_TX_ERROR))
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Model(dao.SysOrder.Table()).
		WhereIn(dao.SysOrder.Columns().UserId, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED, response.CodeMsg(response.DELETE_FAILED))
	}

	_, err = tx.Model(dao.SysRecharge.Table()).
		WhereIn(dao.SysRecharge.Columns().UserId, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED, response.CodeMsg(response.DELETE_FAILED))
	}

	_, err = tx.Model(dao.SysBalance.Table()).
		WhereIn(dao.SysBalance.Columns().UserId, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED, response.CodeMsg(response.DELETE_FAILED))
	}

	_, err = tx.Model(dao.SysUserBill.Table()).
		WhereIn(dao.SysUserBill.Columns().UserId, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED, response.CodeMsg(response.DELETE_FAILED))
	}

	_, err = tx.Model(dao.SysUser.Table()).
		WhereIn(dao.SysUser.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED, response.CodeMsg(response.DELETE_FAILED))
	}

	return
}
