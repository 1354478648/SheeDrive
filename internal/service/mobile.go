package service

import (
	"SheeDrive/internal/model"
	"context"
)

// 1. 定义接口
type IMobile interface {
	GetSwiper(ctx context.Context) (out *model.MobileGetSwiperOutput, err error)
}

// 2. 定义接口变量
var localMobile IMobile

// 3. 定义获取接口实例的函数
func Mobile() IMobile {
	if localMobile == nil {
		panic("接口IMobile未实现或未注册")
	}
	return localMobile
}

// 4. 定义接口实现的注册方法
func RegisterMobile(i IMobile) {
	localMobile = i
}
