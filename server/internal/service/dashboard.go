package service

import (
	"context"
	dao_dashboard "server/internal/type/dashboard/dao"
)

// 定义显示接口
type IDashboard interface {
	GetDetail(ctx context.Context) (res *dao_dashboard.Detail, err error)
}

// 定义接口变量
var localDashboard IDashboard

// 定义一个获取接口的方法
func Dashboard() IDashboard {
	if localDashboard == nil {
		panic("implement not found for interface IDashboard, forgot register?")
	}
	return localDashboard
}

// 定义一个接口实现的注册方法
func RegisterDashboard(i IDashboard) {
	localDashboard = i
}
