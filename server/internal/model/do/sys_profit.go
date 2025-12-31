// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProfit is the golang structure of table sys_profit for DAO operations like Where/Data.
type SysProfit struct {
	g.Meta     `orm:"table:sys_profit, do:true"`
	Id         interface{} //
	Related    interface{} //
	Type       interface{} //
	Amount     interface{} //
	CreateTime *gtime.Time //
}
