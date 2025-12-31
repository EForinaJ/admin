package comment

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/service"
	dto_comment "server/internal/type/cooment/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
)

type sComment struct{}

// Apply implements service.IComment.
func (s *sComment) Apply(ctx context.Context, req *dto_comment.Apply) (err error) {
	_, err = dao.SysComment.Ctx(ctx).Where(dao.SysComment.Columns().Id, req.Id).Data(g.Map{
		dao.SysComment.Columns().Status:   req.Status,
		dao.SysComment.Columns().ManageId: ctx.Value("userId"),
	}).Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}
	return
}

// CheckApply implements service.IComment.
func (s *sComment) CheckApply(ctx context.Context, req *dto_comment.Apply) (err error) {
	exist, err := dao.SysComment.Ctx(ctx).Where(dao.SysComment.Columns().Id, req.Id).
		Where(dao.SysComment.Columns().Status, consts.StatusApply).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if !exist {
		return utils_error.Err(response.FAILD, "评论未找到或已被审核")
	}
	return
}

func init() {
	service.RegisterComment(&sComment{})
}
