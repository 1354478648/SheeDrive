package service

import (
	"SheeDrive/internal/model"
	"SheeDrive/internal/model/entity"
	"context"
)

// 1. 定义接口
type IAdmin interface {
	Login(ctx context.Context, username string, password string) (admin *entity.Admin, err error)
	GetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error)
	GetById(ctx context.Context, id int64) (admin *entity.Admin, err error)
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
