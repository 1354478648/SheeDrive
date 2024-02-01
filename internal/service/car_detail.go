package service

import (
	"SheeDrive/internal/model"
	"context"
)

// 1. 定义接口
type ICarDetail interface {
	GetList(ctx context.Context, in model.CarDetailGetListInput) (out *model.CarDetailGetListOutput, err error)
}

// 2. 定义接口变量
var localCarDetail ICarDetail

// 3. 定义获取接口实例的函数
func CarDetail() ICarDetail {
	if localCarDetail == nil {
		panic("接口ICarDetail未实现或未注册")
	}
	return localCarDetail
}

// 4. 定义接口实现的注册方法
func RegisterCarDetail(i ICarDetail) {
	localCarDetail = i
}
