package witkey

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_witkey "server/internal/type/witkey/dao"
	dto_witkey "server/internal/type/witkey/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetCommissionList implements service.IWitkey.
func (s *sWitkey) GetCommissionList(ctx context.Context, req *dto_witkey.CommissionQuery) (total int, res []*dao_witkey.CommissionList, err error) {
	m := dao.SysCommission.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysCommission.Columns().CreateTime)
	m = m.WhereIn(dao.SysCommission.Columns().WitkeyId, req.Id)
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var list []*entity.SysCommission
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_witkey.CommissionList, len(list))
	for i, v := range list {
		var entity *dao_witkey.CommissionList
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		res[i] = entity
	}

	return
}
