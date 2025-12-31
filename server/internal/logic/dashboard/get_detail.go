package dashboard

import (
	"context"

	"server/internal/consts"
	dao_dashboard "server/internal/type/dashboard/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
)

// GetDetail implements service.IDashboard.
func (s *sDashboard) GetDetail(ctx context.Context) (res *dao_dashboard.Detail, err error) {
	obj, err := g.Redis().Get(ctx, consts.AdminDashboard)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
	}
	if obj.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}
	var detail *dao_dashboard.Detail
	err = obj.Scan(&detail)
	if err != nil {
		return nil, utils_error.Err(response.PARAM_INVALID, response.CodeMsg(response.PARAM_INVALID))
	}

	return detail, nil
}
