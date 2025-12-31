package v1

import (
	dao_comment "server/internal/type/cooment/dao"
	dto_comment "server/internal/type/cooment/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/comment/list" method:"get" tags:"评价" summary:"评价列表"`
	*dto_comment.Query
}
type GetListRes struct {
	Total int                 `json:"total" dc:"总数"`
	List  []*dao_comment.List `json:"list" dc:"评价列表"`
}

type DeleteReq struct {
	g.Meta `path:"/comment/delete" method:"post" tags:"评价" summary:"删除评价"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}

type ApplyReq struct {
	g.Meta `path:"/comment/apply" method:"post" tags:"评价" summary:"审核评价"`
	*dto_comment.Apply
}
type ApplyRes struct{}
