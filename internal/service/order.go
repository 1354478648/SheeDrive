package service

import (
	"SheeDrive/internal/model"
	"context"
)

// 1. 定义接口
type IOrder interface {
	GetList(ctx context.Context, in model.OrderGetListInput) (out *model.OrderGetListOutput, err error)
	GetById(ctx context.Context, in model.OrderGetByIdInput) (out *model.OrderGetByIdOutput, err error)
	Add(ctx context.Context, in model.OrderAddInput) (out *model.OrderAddOutput, err error)
	Delete(ctx context.Context, in model.OrderDeleteInput) (err error)
	UpdateCancel(ctx context.Context, in model.OrderUpdateInput) (err error)
	UpdateConfirm(ctx context.Context, in model.OrderUpdateInput) (err error)
	UpdateSign(ctx context.Context, in model.OrderUpdateInput) (err error)
	UpdateStart(ctx context.Context, in model.OrderUpdateInput) (err error)
	UpdateEnd(ctx context.Context, in model.OrderUpdateInput) (err error)
	UpdateEndAll(ctx context.Context, in model.OrderUpdateInput) (err error)
	GetCarRank(ctx context.Context) (out *model.OrderGetCarRankOutput, err error)
	GetIncomplete(ctx context.Context, in model.OrderGetIncompleteInput) (out *model.OrderGetIncompleteOutput, err error)
	GetAddressTimes(ctx context.Context, in model.OrderGetAddressTimesInput) (out *model.OrderGetAddressTimesOutput, err error)
	GetTimeCount(ctx context.Context, in model.OrderGetTimeCountInput) (out *model.OrderGetTimeCountOutput, err error)
	GetByUserId(ctx context.Context, in model.OrderGetByUserIdInput) (out *model.OrderGetByUserIdOutput, err error)
}

// 2. 定义接口变量
var localOrder IOrder

// 3. 定义获取接口实例的函数
func Order() IOrder {
	if localOrder == nil {
		panic("接口IOrder未实现或未注册")
	}
	return localOrder
}

// 4. 定义接口实现的注册方法
func RegisterOrder(i IOrder) {
	localOrder = i
}
