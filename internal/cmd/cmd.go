package cmd

import (
	"SheeDrive/internal/controller/admin"
	cardetail "SheeDrive/internal/controller/car_detail"
	"SheeDrive/internal/controller/dealer"
	"SheeDrive/internal/controller/file"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/admin", func(group *ghttp.RouterGroup) {
					group.Bind(admin.AdminController)
				})
				group.Group("/dealer", func(group *ghttp.RouterGroup) {
					group.Bind(dealer.DealerController)
				})
				group.Group("/file", func(group *ghttp.RouterGroup) {
					group.Bind(file.FileController)
				})
				group.Group("/cardetail", func(group *ghttp.RouterGroup) {
					group.Bind(cardetail.CarDetailController)
				})
			})
			s.Run()
			return nil
		},
	}
)
