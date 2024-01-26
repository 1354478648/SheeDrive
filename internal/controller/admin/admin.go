package admin

import (
	apiAdmin "SheeDrive/api"
	"SheeDrive/internal/consts"
	"SheeDrive/internal/model"
	"SheeDrive/internal/model/entity"
	"SheeDrive/internal/service"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var AdminController = &cAdmin{}

type cAdmin struct{}

// 生成Token
func jwtToken(admin *entity.Admin) string {
	claim := jwt.RegisteredClaims{
		Subject: admin.Username,
		// 设置过期时间
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(consts.JwtTokenKey))

	if err != nil {
		panic("Token生成错误！")
	}

	return token
}

// 管理员登录
func (c *cAdmin) AdminLogin(ctx context.Context, req *apiAdmin.AdminLoginReq) (res *apiAdmin.AdminLoginRes, err error) {
	// 调用Service层接口
	admin, err := service.Admin().Login(ctx, req.Username, req.Password)
	// 返回Token和管理员基本信息
	if err != nil {
		return nil, err
	}
	res = &apiAdmin.AdminLoginRes{
		Token: jwtToken(admin),
		Admin: &entity.Admin{
			Id:       admin.Id,
			Name:     admin.Name,
			Username: admin.Username,
			Avatar:   admin.Avatar,
			Phone:    admin.Phone,
		},
	}
	return
}

// 管理员查询
func (c *cAdmin) GetAdminList(ctx context.Context, req *apiAdmin.AdminGetListReq) (res *apiAdmin.AdminGetListRes, err error) {
	// 调用service层接口
	adminList, err := service.Admin().GetList(ctx, model.AdminGetListInput{
		Page:       req.CommonPaginationReq.Page,
		PageSize:   req.CommonPaginationReq.Size,
		Username:   req.Username,
		Name:       req.Name,
		BeforeDate: req.BeforeDate,
		AfterDate:  req.AfterDate,
	})

	if err != nil {
		return nil, err
	}
	res = &apiAdmin.AdminGetListRes{
		AdminList: adminList.Items,
		CommonPaginationRes: apiAdmin.CommonPaginationRes{
			Page:  req.Page,
			Size:  req.Size,
			Total: adminList.Total,
		},
	}
	return
}
