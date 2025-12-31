// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysComment is the golang structure for table sys_comment.
type SysComment struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	ManageId   int64       `json:"manageId"   orm:"manage_id"   description:""` //
	UserId     int64       `json:"userId"     orm:"user_id"     description:""` //
	ProductId  int64       `json:"productId"  orm:"product_id"  description:""` //
	Content    string      `json:"content"    orm:"content"     description:""` //
	Rate       int         `json:"rate"       orm:"rate"        description:""` //
	Images     string      `json:"images"     orm:"images"      description:""` //
	IsTop      int         `json:"isTop"      orm:"is_top"      description:""` //
	Status     int         `json:"status"     orm:"status"      description:""` //
	Ip         string      `json:"ip"         orm:"ip"          description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:""` //
}
