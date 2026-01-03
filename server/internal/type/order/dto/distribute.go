package dto_order

type Distribute struct {
	Id       int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
	WitkeyId int64 `p:"witkeyId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
	Type     int   `p:"type" v:"required|in:1,2#请选择结算类型|状态值只能在1或2" dc:"派单类型"`
}
