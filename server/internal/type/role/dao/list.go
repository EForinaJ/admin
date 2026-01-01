package dao_role

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id          int64       `json:"id" dc:"角色ID"`
	Name        string      `json:"name" dc:"角色名称"`
	Code        string      `json:"code" dc:"角色编码"`
	Type        int         `json:"type" dc:"角色类型"`
	Description string      `json:"description" dc:"角色描述"`
	CreateTime  *gtime.Time `json:"createTime" dc:"创建时间"`
}
