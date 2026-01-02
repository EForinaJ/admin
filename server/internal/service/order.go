package service

import (
	"context"
	dao_order "server/internal/type/order/dao"
	dto_order "server/internal/type/order/dto"
)

// 定义显示接口
type IOrder interface {
	GetList(ctx context.Context, req *dto_order.Query) (total int, res []*dao_order.List, err error)
	GetDetail(ctx context.Context, id int64) (res *dao_order.Detail, err error)
	GetDistributeList(ctx context.Context, req *dto_order.DistributeQuery) (total int, res []*dao_order.DistributeList, err error)
	GetLogList(ctx context.Context, req *dto_order.LogQuery) (total int, res []*dao_order.LogList, err error)

	// 创建订单
	Refund(ctx context.Context, req *dto_order.Refund) (err error)
	Delete(ctx context.Context, ids []int64) (err error)
	Paid(ctx context.Context, req *dto_order.Paid) (err error)
	AddDiscount(ctx context.Context, req *dto_order.AddDiscount) (err error)
	Cancel(ctx context.Context, id int64) (err error)
	StartService(ctx context.Context, id int64) (err error)
	Complete(ctx context.Context, id int64) (err error)
	Distribute(ctx context.Context, req *dto_order.Distribute) (err error)
	DistributeCancel(ctx context.Context, req *dto_order.DistributeCancel) (err error)

	CheckComplete(ctx context.Context, id int64) (err error)
	CheckStartService(ctx context.Context, id int64) (err error)
	CheckPaid(ctx context.Context, req *dto_order.Paid) (err error)
	CheckCancel(ctx context.Context, id int64) (err error)
	CheckDiscount(ctx context.Context, req *dto_order.AddDiscount) (err error)
	CheckRefund(ctx context.Context, req *dto_order.Refund) (err error)
	CheckDistribute(ctx context.Context, req *dto_order.Distribute) (err error)
	CheckDistributeCancel(ctx context.Context, req *dto_order.DistributeCancel) (err error)
}

// 定义接口变量
var localOrder IOrder

// 定义一个获取接口的方法
func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

// 定义一个接口实现的注册方法
func RegisterOrder(i IOrder) {
	localOrder = i
}
