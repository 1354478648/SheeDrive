package middleware

import (
	"SheeDrive/internal/consts"
	"SheeDrive/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
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
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(consts.JwtTokenKey), nil
		})

		if err != nil || !token.Valid {
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
