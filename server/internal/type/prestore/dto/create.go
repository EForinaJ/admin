package dto_prestore

type Create struct {
	Amount      float64 `p:"amount"  v:"required#请输入金额"   dc:"金额"`
	BonusAmount float64 `p:"bonusAmount"   dc:"赠送金额"`
	UserId      int64   `p:"userId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
