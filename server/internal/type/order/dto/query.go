package dto_order

type Query struct {
	Page   int    `p:"page"  v:"required|min:1#请输入页数|值最小是1" dc:"列表页数"`
	Limit  int    `p:"limit" v:"required|min:8#请输入条数|值最小是8" dc:"列表条数"`
	Name   string `p:"name" dc:"商品昵称"`
	Code   string `p:"code" dc:"订单号"`
	Status int    `p:"status" dc:"订单状态"`
}

type LogQuery struct {
	Page  int   `p:"page"  v:"required|min:1#请输入页数|值最小是1" dc:"列表页数"`
	Limit int   `p:"limit" v:"required|min:8#请输入条数|值最小是8" dc:"列表条数"`
	Id    int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
