package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"

	"server/internal/lib/casbin"
	_ "server/internal/logic"
	"server/internal/router"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			casbin.InitEnforcer(ctx)
			s := g.Server()
			s.AddStaticPath("/public", "../../public")
			router.LoadRouter(s)
			s.Run()
			return nil
		},
	}
)
