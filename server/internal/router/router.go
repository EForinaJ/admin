package router

import (
	"server/internal/controller/account"
	"server/internal/controller/aftersales"
	"server/internal/controller/auth"
	"server/internal/controller/capital"
	"server/internal/controller/category"
	"server/internal/controller/comment"
	"server/internal/controller/dashboard"
	"server/internal/controller/distribute"
	"server/internal/controller/game"
	"server/internal/controller/level"
	"server/internal/controller/manage"
	"server/internal/controller/media"
	"server/internal/controller/menu"
	"server/internal/controller/order"
	"server/internal/controller/permission"
	"server/internal/controller/prestore"
	"server/internal/controller/product"
	"server/internal/controller/recharge"
	"server/internal/controller/role"
	"server/internal/controller/settlement"
	"server/internal/controller/site"
	"server/internal/controller/system"
	"server/internal/controller/title"
	"server/internal/controller/upload"
	"server/internal/controller/user"
	"server/internal/controller/withdraw"
	"server/internal/controller/witkey"
	"server/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func LoadRouter(s *ghttp.Server) {
	s.Group("/api/v1/admin", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			auth.NewV1(),
			site.NewV1(),
		).Middleware(middleware.Response)
		group.Bind(
			account.NewV1(),
		).Middleware(middleware.Auth).Middleware(middleware.Response)
		group.Middleware(middleware.Auth)
		group.Middleware(middleware.Casbin)
		group.Middleware(middleware.Response)
		group.Bind(
			dashboard.NewV1(),
			menu.NewV1(),
			manage.NewV1(),
			role.NewV1(),
			permission.NewV1(),
			media.NewV1(),
			upload.NewV1(),

			user.NewV1(),
			witkey.NewV1(),
			level.NewV1(),
			game.NewV1(),
			category.NewV1(),
			title.NewV1(),

			product.NewV1(),
			order.NewV1(),
			comment.NewV1(),
			distribute.NewV1(),
			settlement.NewV1(),
			aftersales.NewV1(),
			capital.NewV1(),

			prestore.NewV1(),
			recharge.NewV1(),
			withdraw.NewV1(),
			system.NewV1(),
		)
	})
}
