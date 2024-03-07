package service

import (
	"SheeDrive/internal/model"
	"context"
)

// 1. 定义接口
type IAddress interface {
	GetList(ctx context.Context, in model.UserAddressGetListInput) (out *model.UserAddressGetListOutput, err error)
	Add(ctx context.Context, in model.UserAddressAddInput) (out *model.UserAddressAddOutput, err error)
	GetById(ctx context.Context, in model.UserAddressGetByIdInput) (out *model.UserAddressGetByIdOutput, err error)
	// Update(ctx context.Context, in model.AddressUpdateInput) (err error)
	// Delete(ctx context.Context, in model.AddressDeleteInput) (err error)
}

// 2. 定义接口变量
var localAddress IAddress

// 3. 定义获取接口实例的函数
func Address() IAddress {
	if localAddress == nil {
		panic("接口IAddress未实现或未注册")
	}
	return localAddress
}

// 4. 定义接口实现的注册方法
func RegisterAddress(i IAddress) {
	localAddress = i
}
