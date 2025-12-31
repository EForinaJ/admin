package v1

import (
	dao_account "server/internal/type/account/dao"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDetailReq struct {
	g.Meta `path:"/account" method:"get" tags:"账户" summary:"获取登录账户信息"`
}
type GetDetailRes struct {
	*dao_account.Detail
}

type GetMenuReq struct {
	g.Meta `path:"/account/menus" method:"get" tags:"账户" summary:"获取登录账户路由"`
}
type GetMenuRes struct {
	List []*dao_account.Menu `json:"list" dc:"菜单列表"`
}
