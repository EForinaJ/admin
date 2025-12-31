package dao_dashboard

type Detail struct {
	UserCount      *UserCount      `json:"userCount" dc:"用户数量"`
	OrderCount     *OrderCount     `json:"orderCount" dc:"订单数量"`
	SalesAmount    *SalesAmount    `json:"salesAmount" dc:"销售额度"`
	ProfitAmount   *ProfitAmount   `json:"profitAmount" dc:"收入额度"`
	OrderStatistic *OrderStatistic `json:"orderStatistic" dc:"订单统计"`
	UserStatistic  *UserStatistic  `json:"userStatistic" dc:"用户统计"`

	PendingServiceOrder *PendingServiceOrder `json:"pendingServiceOrder" dc:"待派单服务"`
	ApplySettlement     *ApplySettlement     `json:"applySettlement" dc:"待审核结算报单"`
	ApplyWithdraw       *ApplyWithdraw       `json:"applyWithdraw" dc:"待审核提现"`
	ApplyComment        *ApplyComment        `json:"applyComment" dc:"待审核评论"`
}

type PendingServiceOrder struct {
	TotalCount int `json:"TotalCount" dc:"待派单服务总数"`
	TodayCount int `json:"TodayCount" dc:"今日待服务订单"`
}

type ApplySettlement struct {
	TotalCount int `json:"TotalCount" dc:"待审核结算总数"`
	TodayCount int `json:"TodayCount" dc:"今日待结算"`
}
type ApplyWithdraw struct {
	TotalCount int `json:"TotalCount" dc:"待审核提现总数"`
	TodayCount int `json:"TodayCount" dc:"今日待提现"`
}

type ApplyComment struct {
	TotalCount int `json:"TotalCount" dc:"待审核评论总数"`
	TodayCount int `json:"TodayCount" dc:"今日待审核评论"`
}

type UserCount struct {
	TotalCount int `json:"totalCount" dc:"用户总数"`
	TodayCount int `json:"todayCount" dc:"今日注册"`
}

type OrderCount struct {
	TotalCount int `json:"totalCount" dc:"订单总数"`
	TodayCount int `json:"todayCount" dc:"今日订单"`
}

type SalesAmount struct {
	AmountTotal float64 `json:"amountTotal" dc:"销售总额"`
	AmountToday float64 `json:"amountToday" dc:"今日销售"`
}

type ProfitAmount struct {
	AmountTotal float64 `json:"amountTotal" dc:"收入总额"`
	AmountToday float64 `json:"amountToday" dc:"今日收入"`
}

type OrderStatistic struct {
	Weekdays []string `json:"weekdays" dc:"日期"`
	Count    []int    `json:"count" dc:"数量"`
}
type UserStatistic struct {
	Days  []string `json:"days" dc:"日期"`
	Count []int    `json:"count" dc:"数量"`
}
