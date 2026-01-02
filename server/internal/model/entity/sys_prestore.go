// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPrestore is the golang structure for table sys_prestore.
type SysPrestore struct {
	Id          int64       `json:"id"          orm:"id"           description:""` //
	UserId      int64       `json:"userId"      orm:"user_id"      description:""` //
	ManageId    int64       `json:"manageId"    orm:"manage_id"    description:""` //
	Amount      float64     `json:"amount"      orm:"amount"       description:""` //
	BonusAmount float64     `json:"bonusAmount" orm:"bonus_amount" description:""` //
	Status      int         `json:"status"      orm:"status"       description:""` //
	Reason      string      `json:"reason"      orm:"reason"       description:""` //
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"  description:""` //
}
