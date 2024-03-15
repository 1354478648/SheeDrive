package service

import (
	"SheeDrive/internal/model"
	"context"
)

// 1. 定义接口
type IComment interface {
	GetList(ctx context.Context, in model.CommentGetListInput) (out *model.CommentGetListOutput, err error)
	GetById(ctx context.Context, in model.CommentGetByIdInput) (out *model.CommentGetByIdOutput, err error)
	Add(ctx context.Context, in model.CommentAddInput) (out *model.CommentAddOutput, err error)
	Delete(ctx context.Context, in model.CommentDeleteInput) (err error)
	GetAvg(ctx context.Context, in model.CommentGetAvgInput) (out *model.CommentGetAvgOutput, err error)
	GetByOrderId(ctx context.Context, in model.CommentGetByOrderIdInput) (out *model.CommentGetByOrderIdOutput, err error)
}

// 2. 定义接口变量
var localComment IComment

// 3. 定义获取接口实例的函数
func Comment() IComment {
	if localComment == nil {
		panic("接口IComment未实现或未注册")
	}
	return localComment
}

// 4. 定义接口实现的注册方法
func RegisterComment(i IComment) {
	localComment = i
}
