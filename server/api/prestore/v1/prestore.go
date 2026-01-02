package v1

import (
	dao_prestore "server/internal/type/prestore/dao"
	dto_prestore "server/internal/type/prestore/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type GetInfoReq struct {
	g.Meta `path:"/prestore/detail" method:"get" tags:"预存申请" summary:"预存详情"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetInfoRes struct {
	*dao_prestore.Detail
}

type GetListReq struct {
	g.Meta `path:"/prestore/list" method:"get" tags:"预存申请" summary:"预存申请列表"`
	*dto_prestore.Query
}
type GetListRes struct {
	Total int                  `json:"total" dc:"总数"`
	List  []*dao_prestore.List `json:"list" dc:"预存申请列表"`
}

type DeleteReq struct {
	g.Meta `path:"/prestore/delete" method:"post" tags:"预存申请" summary:"删除预存申请"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}

type ApplyReq struct {
	g.Meta `path:"/prestore/apply" method:"post" tags:"预存申请" summary:"预存审核"`
	*dto_prestore.Apply
}
type ApplyRes struct{}

type CreateReq struct {
	g.Meta `path:"/prestore/create" method:"post" tags:"预存申请" summary:"添加预存"`
	*dto_prestore.Create
}
type CreateRes struct{}
