package dao_comment

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id         int64       `json:"id" dc:"ID"`
	User       string      `json:"user" dc:"评价用户"`
	Product    *Product    `json:"product" dc:"评价用户"`
	Content    string      `json:"content" dc:"评价内容"`
	Images     []string    `json:"images" dc:"评价图片"`
	Status     int         `json:"status" dc:"状态"`
	Rate       int         `json:"type" dc:"评分"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
}
