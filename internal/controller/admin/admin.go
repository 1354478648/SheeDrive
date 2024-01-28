package admin

import (
	apiAdmin "SheeDrive/api"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"SheeDrive/utility"
	"context"
)

var AdminController = &cAdmin{}

type cAdmin struct{}

// 管理员登录
func (c *cAdmin) AdminLogin(ctx context.Context, req *apiAdmin.AdminLoginReq) (res *apiAdmin.AdminLoginRes, err error) {
	// 调用Service层接口
	admin, err := service.Admin().Login(ctx, model.AdminLoginInput{
		Username: req.Username,
		Password: req.Password,
	})
	// 返回Token和管理员基本信息
	if err != nil {
		return nil, err
	}
	res = &apiAdmin.AdminLoginRes{
		Token:     utility.GenToken(admin.Username),
		AdminInfo: admin.AdminBase,
	}
	return
}

// 管理员查询
func (c *cAdmin) GetAdminList(ctx context.Context, req *apiAdmin.AdminGetListReq) (res *apiAdmin.AdminGetListRes, err error) {
	// 调用service层接口
	out, err := service.Admin().GetList(ctx, model.AdminGetListInput{
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
		CommonPaginationRes: apiAdmin.CommonPaginationRes{
			Page:  out.Page,
			Size:  out.PageSize,
			Total: out.Total,
			List:  out.Items,
		},
	}
	return
}
