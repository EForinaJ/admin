// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPrestore is the golang structure of table sys_prestore for DAO operations like Where/Data.
type SysPrestore struct {
	g.Meta      `orm:"table:sys_prestore, do:true"`
	Id          interface{} //
	UserId      interface{} //
	ManageId    interface{} //
	Amount      interface{} //
	BonusAmount interface{} //
	Status      interface{} //
	Reason      interface{} //
	CreateTime  *gtime.Time //
}
