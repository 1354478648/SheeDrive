package cmd

import (
	"SheeDrive/internal/controller/address"
	"SheeDrive/internal/controller/admin"
	cardetail "SheeDrive/internal/controller/car_detail"
	"SheeDrive/internal/controller/dealer"
	"SheeDrive/internal/controller/file"
	"SheeDrive/internal/controller/mobile"
	"SheeDrive/internal/controller/order"
	"SheeDrive/internal/controller/sms"
	"SheeDrive/internal/controller/stock"
	"SheeDrive/internal/controller/user"
	"SheeDrive/internal/service"
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
				// 无需鉴权的路由组
				group.Group("/admin", func(group *ghttp.RouterGroup) {
					group.Bind(admin.AdminController.AdminLogin)
					group.Bind(admin.AdminController.AdminUpdatePasswordByPhone)
				})
				group.Group("/dealer", func(group *ghttp.RouterGroup) {
					group.Bind(dealer.DealerController.DealerLogin)
					group.Bind(dealer.DealerController.DealerUpdatePasswordByPhone)
				})
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Bind(user.UserController.UserLogin)
					group.Bind(user.UserController.UserRegister)
					group.Bind(user.UserController.UserLoginByPhone)
					group.Bind(user.UserController.UserUpdatePasswordByPhone)
				})
				group.Group("/mobile", func(group *ghttp.RouterGroup) {
					group.Bind(mobile.MobileController)
				})
				group.Group("/sms", func(group *ghttp.RouterGroup) {
					group.Bind(sms.SmsController)
				})
				// 需要鉴权的路由组
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Group("/admin", func(group *ghttp.RouterGroup) {
						group.Bind(
							admin.AdminController.AdminAdd,
							admin.AdminController.AdminDelete,
							admin.AdminController.AdminGetById,
							admin.AdminController.AdminResetPassword,
							admin.AdminController.AdminUpdate,
							admin.AdminController.AdminUpdateAvatar,
							admin.AdminController.AdminUpdatePassword,
							admin.AdminController.AdminUpdateStatus,
							admin.AdminController.GetAdminList,
						)
					})
					group.Group("/dealer", func(group *ghttp.RouterGroup) {
						group.Bind(
							dealer.DealerController.DealerAdd,
							dealer.DealerController.DealerDelete,
							dealer.DealerController.DealerGetById,
							dealer.DealerController.DealerList,
							dealer.DealerController.DealerResetPassword,
							dealer.DealerController.DealerUpdate,
							dealer.DealerController.DealerUpdateAvatar,
							dealer.DealerController.DealerUpdatePassword,
							dealer.DealerController.DealerUpdateStatus,
						)
					})
					group.Group("/user", func(group *ghttp.RouterGroup) {
						group.Bind(
							user.UserController.UserDelete,
							user.UserController.UserGetById,
							user.UserController.UserGetList,
							user.UserController.UserUpdateAvatar,
							user.UserController.UserUpdatePassword,
							user.UserController.UserUpdateStatus,
						)
					})
					group.Group("/file", func(group *ghttp.RouterGroup) {
						group.Bind(file.FileController)
					})
					group.Group("/cardetail", func(group *ghttp.RouterGroup) {
						group.Bind(cardetail.CarDetailController)
					})
					group.Group("/stock", func(group *ghttp.RouterGroup) {
						group.Bind(stock.StockController)
					})
					group.Group("/address", func(group *ghttp.RouterGroup) {
						group.Bind(address.AddressController)
					})
					group.Group("/order", func(group *ghttp.RouterGroup) {
						group.Bind(order.OrderController)
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
