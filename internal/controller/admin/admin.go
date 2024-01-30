package admin

import (
	apiAdmin "SheeDrive/api/admin"
	apiPagination "SheeDrive/api/pagination"
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
		AdminInfo: admin.AdminInfoBase,
	}
	return
}

// 管理员列表查询
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
		CommonPaginationRes: apiPagination.CommonPaginationRes{
			Page:  out.Page,
			Size:  out.PageSize,
			Total: out.Total,
			List:  out.Items,
		},
	}
	return
}

// 通过Id查询管理员
func (c *cAdmin) AdminGetById(ctx context.Context, req *apiAdmin.AdminGetByIdReq) (res *apiAdmin.AdminGetByIdRes, err error) {
	// 调用service层接口
	out, err := service.Admin().GetById(ctx, model.AdminGetByIdInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res = &apiAdmin.AdminGetByIdRes{
		AdminInfo: out.AdminInfoBase,
	}
	return
}

// 管理员添加
func (c *cAdmin) AdminAdd(ctx context.Context, req *apiAdmin.AdminAddReq) (res *apiAdmin.AdminAddRes, err error) {
	// 调用service层接口
	out, err := service.Admin().Add(ctx, model.AdminAddInput{
		AdminAddUpdateBase: model.AdminAddUpdateBase{
			Name:     req.Name,
			Username: req.Username,
			Phone:    req.Phone,
		},
	})
	if err != nil {
		return nil, err
	}
	res = &apiAdmin.AdminAddRes{
		Id: out.Id,
	}
	return
}

// 管理员修改
func (c *cAdmin) AdminUpdate(ctx context.Context, req *apiAdmin.AdminUpdateReq) (res *apiAdmin.AdminUpdateRes, err error) {
	// 调用service层接口
	err = service.Admin().Update(ctx, model.AdminUpdateInput{
		Id: req.Id,
		AdminAddUpdateBase: model.AdminAddUpdateBase{
			Name:     req.Name,
			Username: req.Username,
			Phone:    req.Phone,
		},
	})
	if err != nil {
		return nil, err
	}
	return
}

// 管理员删除
func (c *cAdmin) AdminDelete(ctx context.Context, req *apiAdmin.AdminDeleteReq) (res *apiAdmin.AdminDeleteRes, err error) {
	// 调用service接口
	err = service.Admin().Delete(ctx, model.AdminDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}

// 管理员状态修改
func (c *cAdmin) AdminUpdateStatus(ctx context.Context, req *apiAdmin.AdminUpdateStatusReq) (res *apiAdmin.AdminUpdateStatusRes, err error) {
	// 调用service接口
	err = service.Admin().UpdateStatus(ctx, model.AdminUpdateStatusInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}

// 管理员密码修改
func (c *cAdmin) AdminUpdatePassword(ctx context.Context, req *apiAdmin.AdminUpdatePasswordReq) (res *apiAdmin.AdminUpdatePasswordRes, err error) {
	// 调用service接口
	err = service.Admin().UpdatePassword(ctx, model.AdminUpdatePasswordInput{
		Id:          req.Id,
		OldPassword: req.Password,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return nil, err
	}

	return
}

// 管理员密码重置
func (c *cAdmin) AdminResetPassword(ctx context.Context, req *apiAdmin.AdminResetPasswordReq) (res *apiAdmin.AdminResetPasswordRes, err error) {
	// 调用service接口
	err = service.Admin().ResetPassword(ctx, model.AdminResetPasswordInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return
}
