package service

import (
	"SheeDrive/internal/model"
	"context"
)

// 1. 定义接口
type IDealer interface {
	Login(ctx context.Context, in model.DealerLoginInput) (out *model.DealerLoginOutput, err error)
	GetList(ctx context.Context, in model.DealerGetListInput) (out *model.DealerGetListOutput, err error)
	GetById(ctx context.Context, in model.DealerGetByIdInput) (out *model.DealerGetByIdOutput, err error)
	Add(ctx context.Context, in model.DealerAddInput) (out *model.DealerAddOutput, err error)
	Update(ctx context.Context, in model.DealerUpdateInput) (err error)
	Delete(ctx context.Context, in model.DealerDeleteInput) (err error)
	UpdateStatus(ctx context.Context, in model.DealerUpdateStatusInput) (err error)
	UpdatePassword(ctx context.Context, in model.DealerUpdatePasswordInput) (err error)
	ResetPassword(ctx context.Context, in model.DealerResetPasswordInput) (err error)
	UpdateAvatar(ctx context.Context, in model.DealerUpdateAvatarInput) (err error)
	UpdatePasswordByPhone(ctx context.Context, in model.DealerUpdatePasswordByPhoneInput) (err error)
}

// 2. 定义接口变量
var localDealer IDealer

// 3. 定义获取接口实例的函数
func Dealer() IDealer {
	if localDealer == nil {
		panic("接口IDealer未实现或未注册")
	}
	return localDealer
}

// 4. 定义接口实现的注册方法
func RegisterDealer(i IDealer) {
	localDealer = i
}
