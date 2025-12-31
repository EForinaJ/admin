package dashboard

import (
	"context"

	"server/internal/consts"
	"server/internal/dao"
	dao_dashboard "server/internal/type/dashboard/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

func setData() (err error) {
	ctx := gctx.New()
	_, _ = gcron.Add(ctx, "# */10 * * * *", func(ctx context.Context) {

		var userCount dao_dashboard.UserCount
		userTotalCount, _ := dao.SysUser.Ctx(ctx).Count()
		userCount.TotalCount = userTotalCount
		userTodayCount, _ := dao.SysUser.Ctx(ctx).
			Where("DATE(create_time) = ?", gtime.Now().Format("Y-m-d")).
			Count()
		userCount.TodayCount = userTodayCount

		var orderCount dao_dashboard.OrderCount
		orderTotalCount, _ := dao.SysOrder.Ctx(ctx).Count()
		orderCount.TotalCount = orderTotalCount
		orderTodayCount, _ := dao.SysOrder.Ctx(ctx).
			Where("DATE(create_time) = ?", gtime.Now().Format("Y-m-d")).
			Count()
		orderCount.TodayCount = orderTodayCount

		var salesAmount dao_dashboard.SalesAmount
		salesAmountTotal, _ := dao.SysCapital.Ctx(ctx).
			Where(dao.SysCapital.Columns().Type, consts.CapitalPaymentOrder).
			Sum("amount")

		salesAmount.AmountTotal = salesAmountTotal
		salesAmountToday, _ := dao.SysCapital.Ctx(ctx).
			Where(dao.SysCapital.Columns().Type, consts.CapitalPaymentOrder).
			Where("DATE(create_time) = ?", gtime.Now().Format("Y-m-d")).
			Sum("amount")
		salesAmount.AmountToday = salesAmountToday

		var profitAmount dao_dashboard.ProfitAmount
		profitAmountTotal, _ := dao.SysProfit.Ctx(ctx).
			Sum("amount")
		profitAmount.AmountTotal = profitAmountTotal
		profitAmountToday, _ := dao.SysProfit.Ctx(ctx).
			Where("DATE(create_time) = ?", gtime.Now().Format("Y-m-d")).
			Sum("amount")
		profitAmount.AmountToday = profitAmountToday

		var detail dao_dashboard.Detail
		detail.UserCount = &userCount
		detail.OrderCount = &orderCount
		detail.SalesAmount = &salesAmount
		detail.ProfitAmount = &profitAmount

		weekdays := []string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}
		var orderStatistic dao_dashboard.OrderStatistic
		for i := 0; i < 7; i++ {
			// 计算日期
			targetDate := gtime.Now().AddDate(0, 0, -i)

			// 获取星期（time.Weekday 0=周日, 1=周一...）
			weekday := targetDate.Weekday()
			orderStatistic.Weekdays = append(orderStatistic.Weekdays, weekdays[weekday])
			count, _ := dao.SysOrder.Ctx(ctx).
				Where("DATE(create_time) = ?", targetDate.Format("Y-m-d")).
				Count()
			orderStatistic.Count = append(orderStatistic.Count, count)
		}
		detail.OrderStatistic = &orderStatistic

		var userStatistic dao_dashboard.UserStatistic
		for i := 0; i < 15; i++ {
			// 计算日期
			targetDate := gtime.Now().AddDate(0, 0, -i)
			userStatistic.Days = append(userStatistic.Days, targetDate.Format("Y-m-d"))
			count, _ := dao.SysUser.Ctx(ctx).
				Where("DATE(create_time) = ?", targetDate.Format("Y-m-d")).
				Count()
			userStatistic.Count = append(userStatistic.Count, count)
		}
		detail.UserStatistic = &userStatistic

		var pendingServiceOrder dao_dashboard.PendingServiceOrder
		orderPendingServiceOrderTotalCount, _ := dao.SysOrder.Ctx(ctx).
			Where(dao.SysOrder.Columns().Status, consts.OrderStatusPendingOrder).
			Count()
		pendingServiceOrder.TotalCount = orderPendingServiceOrderTotalCount
		orderPendingServiceOrderTodayCount, _ := dao.SysOrder.Ctx(ctx).
			Where(dao.SysOrder.Columns().Status, consts.OrderStatusPendingOrder).
			Where("DATE(create_time) = ?", gtime.Now().Format("Y-m-d")).
			Count()
		pendingServiceOrder.TodayCount = orderPendingServiceOrderTodayCount
		detail.PendingServiceOrder = &pendingServiceOrder

		var applySettlement dao_dashboard.ApplySettlement
		applySettlementTotalCount, _ := dao.SysSettlement.Ctx(ctx).
			Where(dao.SysSettlement.Columns().Status, consts.StatusApply).
			Count()
		applySettlement.TotalCount = applySettlementTotalCount
		applySettlementTodayCount, _ := dao.SysSettlement.Ctx(ctx).
			Where(dao.SysSettlement.Columns().Status, consts.StatusApply).
			Where("DATE(create_time) = ?", gtime.Now().Format("Y-m-d")).
			Count()
		applySettlement.TodayCount = applySettlementTodayCount
		detail.ApplySettlement = &applySettlement

		var applyWithdraw dao_dashboard.ApplyWithdraw
		applyWithdrawTotalCount, _ := dao.SysWithdraw.Ctx(ctx).
			Where(dao.SysWithdraw.Columns().Status, consts.StatusApply).
			Count()
		applyWithdraw.TotalCount = applyWithdrawTotalCount
		applyWithdrawTodayCount, _ := dao.SysWithdraw.Ctx(ctx).
			Where(dao.SysWithdraw.Columns().Status, consts.StatusApply).
			Where("DATE(create_time) = ?", gtime.Now().Format("Y-m-d")).
			Count()
		applyWithdraw.TodayCount = applyWithdrawTodayCount
		detail.ApplyWithdraw = &applyWithdraw

		var applyComment dao_dashboard.ApplyComment
		applyCommentTotalCount, _ := dao.SysComment.Ctx(ctx).
			Where(dao.SysComment.Columns().Status, consts.StatusApply).
			Count()
		applyComment.TotalCount = applyCommentTotalCount
		applyCommentTodayCount, _ := dao.SysComment.Ctx(ctx).
			Where(dao.SysComment.Columns().Status, consts.StatusApply).
			Where("DATE(create_time) = ?", gtime.Now().Format("Y-m-d")).
			Count()
		applyComment.TodayCount = applyCommentTodayCount
		detail.ApplyComment = &applyComment

		g.Redis().Set(ctx, consts.AdminDashboard, detail)
	}, "dashboard")
	return
}
