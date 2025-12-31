// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProfit is the golang structure for table sys_profit.
type SysProfit struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	Related    string      `json:"related"    orm:"related"     description:""` //
	Type       int         `json:"type"       orm:"type"        description:""` //
	Amount     float64     `json:"amount"     orm:"amount"      description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
