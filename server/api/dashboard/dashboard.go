// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dashboard

import (
	"context"

	"server/api/dashboard/v1"
)

type IDashboardV1 interface {
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
}
