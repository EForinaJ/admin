package dto_role

type Create struct {
	Type        int    `p:"type" v:"required|in:1,2#请输入角色类型|值只能在1和2" dc:"角色类型"`
	Name        string `p:"name" v:"required#请输入角色名称" dc:"角色名称"`
	Code        string `p:"code" v:"required#请输入角色编码" dc:"角色编码"`
	Description string `p:"description" dc:"角色描述"`
}
