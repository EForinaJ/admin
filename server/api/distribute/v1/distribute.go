package v1

import (
	dao_distribute "server/internal/type/distribute/dao"
	dto_distribute "server/internal/type/distribute/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/distribute/list" method:"get" tags:"派单" summary:"派单列表"`
	*dto_distribute.Query
}
type GetListRes struct {
	Total int                    `json:"total" dc:"总数"`
	List  []*dao_distribute.List `json:"list" dc:"派单列表"`
}

type CancelReq struct {
	g.Meta `path:"/distribute/cancel" method:"post" tags:"派单" summary:"派单取消"`
	*dto_distribute.Cancel
}
type CancelRes struct{}

type GetDetailReq struct {
	g.Meta `path:"/distribute/detail" method:"get" tags:"派单" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_distribute.Detail
}
