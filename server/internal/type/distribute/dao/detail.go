package dao_distribute

import (
	dao_order "server/internal/type/order/dao"

	"github.com/gogf/gf/v2/os/gtime"
)

type Detail struct {
	Id         int64             `json:"id" dc:"项目Id"`
	Order      *dao_order.Detail `json:"order" dc:"订单信息"`
	Manage     string            `json:"manage" dc:"派单客服"`
	Witkey     string            `json:"witkey" dc:"接单威客"`
	Game       string            `json:"game" dc:"游戏领域"`
	Title      string            `json:"title" dc:"头衔勋章"`
	Type       int               `json:"type" dc:"类型"`
	Status     int               `json:"status" dc:"状态"`
	Reason     string            `json:"reason" dc:"取消原因"`
	CreateTime *gtime.Time       `json:"createTime" dc:"下单时间"`
}
