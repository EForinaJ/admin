package dto_account

type Edit struct {
	Email       string   `p:"email" v:"email#邮箱格式不正确"     dc:"邮箱"`
	Avatar      string   `p:"avatar" dc:"头像"`
	Name        string   `p:"name" dc:"名称"`
	Address     string   `p:"address"  dc:"地址"`
	Description string   `p:"description"  dc:"描述"`
	Sex         int      `p:"sex"  dc:"性别"`
	Tags        []string `p:"tags"  dc:"标签"`
}
