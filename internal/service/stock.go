package service

import (
	"SheeDrive/internal/model"
	"context"
)

// 1. 定义接口
type IStock interface {
	GetList(ctx context.Context, in model.StockGetListInput) (out *model.StockGetListOutput, err error)
	GetById(ctx context.Context, in model.StockGetByIdInput) (out *model.StockGetByIdOutput, err error)
	Add(ctx context.Context, in model.StockAddInput) (out *model.StockAddOutput, err error)
	Delete(ctx context.Context, in model.StockDeleteInput) (err error)
	GetByCarId(ctx context.Context, in model.StockGetByCarIdInput) (out *model.StockGetByCarIdOutput, err error)
}

// 2. 定义接口变量
var localStock IStock

// 3. 定义获取接口实例的函数
func Stock() IStock {
	if localStock == nil {
		panic("接口IStock未实现或未注册")
	}
	return localStock
}

// 4. 定义接口实现的注册方法
func RegisterStock(i IStock) {
	localStock = i
}
