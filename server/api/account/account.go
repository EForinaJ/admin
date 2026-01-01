// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package account

import (
	"context"

	"server/api/account/v1"
)

type IAccountV1 interface {
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	GetMenu(ctx context.Context, req *v1.GetMenuReq) (res *v1.GetMenuRes, err error)
	Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error)
	ChangePass(ctx context.Context, req *v1.ChangePassReq) (res *v1.ChangePassRes, err error)
}
