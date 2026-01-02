package v1

import (
	dao_order "server/internal/type/order/dao"
	dto_order "server/internal/type/order/dto"
	dao_witkey "server/internal/type/witkey/dao"
	dto_witkey "server/internal/type/witkey/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/order/list" method:"get" tags:"订单" summary:"订单列表"`
	*dto_order.Query
}
type GetListRes struct {
	Total int               `json:"total" dc:"总数"`
	List  []*dao_order.List `json:"list" dc:"订单列表"`
}

type GetDetailReq struct {
	g.Meta `path:"/order/detail" method:"get" tags:"订单" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_order.Detail
}

type RefundReq struct {
	g.Meta `path:"/order/refund" method:"post" tags:"订单" summary:"订单退款"`
	*dto_order.Refund
}
type RefundRes struct{}

type AddDiscountReq struct {
	g.Meta `path:"/order/add/discount" method:"post" tags:"订单" summary:"订单添加折扣"`
	*dto_order.AddDiscount
}
type AddDiscountRes struct{}

type DeleteReq struct {
	g.Meta `path:"/order/delete" method:"post" tags:"订单" summary:"删除订单"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}

type PaidReq struct {
	g.Meta `path:"/order/paid" method:"post" tags:"订单" summary:"订单确认收款"`
	*dto_order.Paid
}
type PaidRes struct{}

type CancelReq struct {
	g.Meta `path:"/order/cancel" method:"post" tags:"订单" summary:"关闭订单"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type CancelRes struct{}

type StartServiceReq struct {
	g.Meta `path:"/order/start" method:"post" tags:"订单" summary:"开始服务"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type StartServiceRes struct{}

type CompleteReq struct {
	g.Meta `path:"/order/complete" method:"post" tags:"订单" summary:"完成服务"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type CompleteRes struct{}

type GetWitkeyListReq struct {
	g.Meta `path:"/order/witkey/list" method:"get" tags:"订单" summary:"威客列表"`
	*dto_witkey.Query
}
type GetWitkeyListRes struct {
	Total int                `json:"total" dc:"总数"`
	List  []*dao_witkey.List `json:"list" dc:"威客列表"`
}

type DistributeReq struct {
	g.Meta `path:"/order/distribute" method:"post" tags:"订单" summary:"派发威客"`
	*dto_order.Distribute
}
type DistributeRes struct{}

type GetDistributeListReq struct {
	g.Meta `path:"/order/distribute/list" method:"get" tags:"订单" summary:"派单列表"`
	*dto_order.DistributeQuery
}
type GetDistributeListRes struct {
	Total int                         `json:"total" dc:"总数"`
	List  []*dao_order.DistributeList `json:"list" dc:"派单列表"`
}

type DistributeCancelReq struct {
	g.Meta `path:"/order/distribute/cancel" method:"post" tags:"订单" summary:"派发取消"`
	*dto_order.DistributeCancel
}
type DistributeCancelRes struct{}

type GetLogListReq struct {
	g.Meta `path:"/order/log/list" method:"get" tags:"订单" summary:"订单日志"`
	*dto_order.LogQuery
}
type GetLogListRes struct {
	Total int                  `json:"total" dc:"总数"`
	List  []*dao_order.LogList `json:"list" dc:"订单日志列表"`
}
