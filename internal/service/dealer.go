package service

import (
	"SheeDrive/internal/model"
	"context"
)

// 1. 定义接口
type IDealer interface {
	Login(ctx context.Context, in model.DealerLoginInput) (out *model.DealerLoginOutput, err error)
	GetList(ctx context.Context, in model.DealerGetListInput) (out *model.DealerGetListOutput, err error)
	// Add(ctx context.Context, in model.DealerAddInput) (out *model.DealerAddOutput, err error)
	// Update(ctx context.Context, in model.DealerUpdateInput) (err error)
	// Delete(ctx context.Context, in model.DealerDeleteInput) (err error)
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
