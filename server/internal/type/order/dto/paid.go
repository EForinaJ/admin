package dto_order

type Paid struct {
	Id      int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
	PayMode int   `p:"payMode" v:"required|in:3,4#请选择结算类型|状态值只能在3或4" dc:"收款类型"`
}
