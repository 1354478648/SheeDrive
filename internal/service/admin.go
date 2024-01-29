package service

import (
	"SheeDrive/internal/model"
	"context"
)

// 1. 定义接口
type IAdmin interface {
	Login(ctx context.Context, in model.AdminLoginInput) (out *model.AdminLoginOutput, err error)
	GetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error)
	GetById(ctx context.Context, in model.AdminGetByIdInput) (out *model.AdminGetByIdOutput, err error)
	Add(ctx context.Context, in model.AdminAddInput) (out *model.AdminAddOutput, err error)
	Update(ctx context.Context, in model.AdminUpdateInput) (err error)
	Delete(ctx context.Context, in model.AdminDeleteInput) (err error)
	UpdateStatus(ctx context.Context, in model.AdminUpdateStatusInput) (err error)
	UpdatePassword(ctx context.Context, in model.AdminUpdatePasswordInput) (err error)
	ResetPassword(ctx context.Context, in model.AdminResetPasswordInput) (err error)
}

// 2. 定义接口变量
var localAdmin IAdmin

// 3. 定义获取接口实例的函数
func Admin() IAdmin {
	if localAdmin == nil {
		panic("接口IAdmin未实现或未注册")
	}
	return localAdmin
}

// 4. 定义接口实现的注册方法
func RegisterAdmin(i IAdmin) {
	localAdmin = i
}
