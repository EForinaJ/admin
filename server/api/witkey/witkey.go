// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package witkey

import (
	"context"

	"server/api/witkey/v1"
)

type IWitkeyV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	GetEidt(ctx context.Context, req *v1.GetEidtReq) (res *v1.GetEidtRes, err error)
	Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error)
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	ChangeCommission(ctx context.Context, req *v1.ChangeCommissionReq) (res *v1.ChangeCommissionRes, err error)
	GetCommissionList(ctx context.Context, req *v1.GetCommissionListReq) (res *v1.GetCommissionListRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}
