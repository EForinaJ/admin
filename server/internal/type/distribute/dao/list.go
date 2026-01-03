package dao_distribute

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"项目Id"`
	Order      string      `json:"order" dc:"订单号"`
	Manage     string      `json:"manage" dc:"派单客服"`
	Witkey     string      `json:"witkey" dc:"接单威客"`
	Game       string      `json:"game" dc:"游戏领域"`
	Title      string      `json:"title" dc:"头衔勋章"`
	Type       int         `json:"type" dc:"类型"`
	Status     int         `json:"status" dc:"状态"`
	CreateTime *gtime.Time `json:"createTime" dc:"下单时间"`
}
