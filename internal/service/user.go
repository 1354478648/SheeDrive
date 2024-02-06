package service

import (
	"SheeDrive/internal/model"
	"context"
)

// 1. 定义接口
type IUser interface {
	Login(ctx context.Context, in model.UserLoginInput) (out *model.UserLoginOutput, err error)
	Register(ctx context.Context, in model.UserRegisterInput) (out *model.UserRegisterOutput, err error)
	GetList(ctx context.Context, in model.UserGetListInput) (out *model.UserGetListOutput, err error)
	GetById(ctx context.Context, in model.UserGetByIdInput) (out *model.UserGetByIdOutput, err error)
	Delete(ctx context.Context, in model.UserDeleteInput) (err error)
	UpdateStatus(ctx context.Context, in model.UserUpdateStatusInput) (err error)
	UpdatePassword(ctx context.Context, in model.UserUpdatePasswordInput) (err error)
	UpdateAvatar(ctx context.Context, in model.UserUpdateAvatarInput) (err error)
}

// 2. 定义接口变量
var localUser IUser

// 3. 定义获取接口实例的函数
func User() IUser {
	if localUser == nil {
		panic("接口IUser未实现或未注册")
	}
	return localUser
}

// 4. 定义接口实现的注册方法
func RegisterUser(i IUser) {
	localUser = i
}
