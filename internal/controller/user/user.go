package user

import (
	apiPagination "SheeDrive/api/pagination"
	apiUser "SheeDrive/api/user"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"SheeDrive/utility"
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
		Token:    utility.GenToken(user.Username),
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
