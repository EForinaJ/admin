// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package order

import (
	"context"

	"server/api/order/v1"
)

type IOrderV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	Refund(ctx context.Context, req *v1.RefundReq) (res *v1.RefundRes, err error)
	AddDiscount(ctx context.Context, req *v1.AddDiscountReq) (res *v1.AddDiscountRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Paid(ctx context.Context, req *v1.PaidReq) (res *v1.PaidRes, err error)
	Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error)
	StartService(ctx context.Context, req *v1.StartServiceReq) (res *v1.StartServiceRes, err error)
	Complete(ctx context.Context, req *v1.CompleteReq) (res *v1.CompleteRes, err error)
	GetWitkeyList(ctx context.Context, req *v1.GetWitkeyListReq) (res *v1.GetWitkeyListRes, err error)
	Distribute(ctx context.Context, req *v1.DistributeReq) (res *v1.DistributeRes, err error)
	GetDistributeList(ctx context.Context, req *v1.GetDistributeListReq) (res *v1.GetDistributeListRes, err error)
	DistributeCancel(ctx context.Context, req *v1.DistributeCancelReq) (res *v1.DistributeCancelRes, err error)
	GetLogList(ctx context.Context, req *v1.GetLogListReq) (res *v1.GetLogListRes, err error)
}
