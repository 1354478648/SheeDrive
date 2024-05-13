package middleware

import (
	"SheeDrive/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type iMiddleware struct{}

func New() *iMiddleware { return &iMiddleware{} }

func init() {
	service.RegisterMiddleware(New())
}

// Auth implements service.IMiddleware.
func (*iMiddleware) Auth(r *ghttp.Request) {
	var res *ghttp.DefaultHandlerResponse
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		res = &ghttp.DefaultHandlerResponse{
			Code:    403,
			Message: "请先登录",
		}
	} else {
		value, err := g.Redis().Get(r.Context(), tokenStr)
		if err != nil || value == nil || value.String() != tokenStr {
			res = &ghttp.DefaultHandlerResponse{
				Code:    403,
				Message: "身份验证已过期，请重新登录",
			}
		}
	}
	if res != nil {
		r.Response.WriteJsonExit(res)
	}
	r.Middleware.Next()
}
