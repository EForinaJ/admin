package v1

import (
	dao_dashboard "server/internal/type/dashboard/dao"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDetailReq struct {
	g.Meta `path:"/dashboard/detail" method:"get" tags:"仪表盘" summary:"获取信息"`
}
type GetDetailRes struct {
	*dao_dashboard.Detail
}
