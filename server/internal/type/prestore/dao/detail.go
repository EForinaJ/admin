package dao_prestore

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id          int64       `json:"id" dc:"ID"`
	Manage      string      `json:"manage" dc:"管理昵称"`
	User        string      `json:"user" dc:"预存用户"`
	Amount      float64     `json:"amount" dc:"预存金额"`
	BonusAmount float64     `json:"bonusAmount" dc:"赠送金额"`
	Reason      string      `json:"reason" dc:"拒绝原因"`
	Status      int         `json:"status" dc:"状态"`
	CreateTime  *gtime.Time `json:"createTime" dc:"创建时间"`
}
