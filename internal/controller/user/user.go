package user

import (
	apiPagination "SheeDrive/api/pagination"
	apiUser "SheeDrive/api/user"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"context"
)

var UserController = &cUser{}

type cUser struct{}

// 用户登录
func (c *cUser) UserLogin(ctx context.Context, req *apiUser.UserLoginReq) (res *apiUser.UserLoginRes, err error) {
	user, err := service.User().Login(ctx, model.UserLoginInput{
		Username: req.UserName,
		Password: req.Password,
	})
	// 返回Token和用户基本信息
	if err != nil {
		return nil, err
	}
	res = &apiUser.UserLoginRes{
		Token:    user.Token,
		UserInfo: user.UserInfoBase,
	}
	return
}

// 用户通过手机号登录
func (c *cUser) UserLoginByPhone(ctx context.Context, req *apiUser.UserLoginByPhoneReq) (res *apiUser.UserLoginByPhoneRes, err error) {
	user, err := service.User().LoginByPhone(ctx, model.UserLoginByPhoneInput{
		Phone: req.Phone,
		Code:  req.Code,
	})
	if err != nil {
		return nil, err
	}
	res = &apiUser.UserLoginByPhoneRes{
		Token:    user.Token,
		UserInfo: user.UserInfoBase,
	}
	return
}

// 用户注册
func (c *cUser) UserRegister(ctx context.Context, req *apiUser.UserRegisterReq) (res *apiUser.UserRegisterRes, err error) {
	out, err := service.User().Register(ctx, model.UserRegisterInput{
		LastName:  req.LastName,
		FirstName: req.FirstName,
		IdNumber:  req.IdNumber,
		Phone:     req.Phone,
		Password:  req.Password,
	})
	if err != nil {
		return nil, err
	}
	res = &apiUser.UserRegisterRes{
		UserInfo: out.UserInfoBase,
	}
	return
}

// 用户列表查询
func (c *cUser) UserGetList(ctx context.Context, req *apiUser.UserGetListReq) (res *apiUser.UserGetListRes, err error) {
	out, err := service.User().GetList(ctx, model.UserGetListInput{
		Page:       req.CommonPaginationReq.Page,
		PageSize:   req.CommonPaginationReq.Size,
		Username:   req.Username,
		Name:       req.Name,
		Status:     req.Status,
		BeforeDate: req.BeforeDate,
		AfterDate:  req.AfterDate,
	})
	if err != nil {
		return nil, err
	}
	res = &apiUser.UserGetListRes{
		CommonPaginationRes: apiPagination.CommonPaginationRes{
			Page:  out.Page,
			Size:  out.PageSize,
			Total: out.Total,
			List:  out.Items,
		},
	}
	return
}

// 通过Id查询用户
func (c *cUser) UserGetById(ctx context.Context, req *apiUser.UserGetByIdReq) (res *apiUser.UserGetByIdRes, err error) {
	out, err := service.User().GetById(ctx, model.UserGetByIdInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res = &apiUser.UserGetByIdRes{
		UserInfo: out.UserInfoBase,
	}
	return
}

// 删除用户
func (c *cUser) UserDelete(ctx context.Context, req *apiUser.UserDeleteReq) (res *apiUser.UserDeleteRes, err error) {
	err = service.User().Delete(ctx, model.UserDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}

// 用户状态修改
func (c *cUser) UserUpdateStatus(ctx context.Context, req *apiUser.UserUpdateStatusReq) (res *apiUser.UserUpdateStatusRes, err error) {
	err = service.User().UpdateStatus(ctx, model.UserUpdateStatusInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}

// 用户密码修改
func (c *cUser) UserUpdatePassword(ctx context.Context, req *apiUser.UserUpdatePasswordReq) (res *apiUser.UserUpdatePasswordRes, err error) {
	err = service.User().UpdatePassword(ctx, model.UserUpdatePasswordInput{
		Id:          req.Id,
		OldPassword: req.Password,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return nil, err
	}
	return
}

// 用户头像修改
func (c *cUser) UserUpdateAvatar(ctx context.Context, req *apiUser.UserUpdateAvatarReq) (res *apiUser.UserUpdateAvatarRes, err error) {
	err = service.User().UpdateAvatar(ctx, model.UserUpdateAvatarInput{
		Id:  req.Id,
		Url: req.Url,
	})
	if err != nil {
		return nil, err
	}
	return
}
