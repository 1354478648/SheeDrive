package dealer

import (
	apiDealer "SheeDrive/api/dealer"
	apiPagination "SheeDrive/api/pagination"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"context"
)

var DealerController = &cDealer{}

type cDealer struct{}

// 经销商登录
func (c *cDealer) DealerLogin(ctx context.Context, req *apiDealer.DealerLoginReq) (res *apiDealer.DealerLoginRes, err error) {
	// 调用Service层接口
	dealer, err := service.Dealer().Login(ctx, model.DealerLoginInput{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	res = &apiDealer.DealerLoginRes{
		Token:      dealer.Token,
		DealerInfo: dealer.DealerInfoBase,
	}

	return
}

// 经销商分页关键字查询
func (c *cDealer) DealerList(ctx context.Context, req *apiDealer.DealerGetListReq) (res *apiDealer.DealerGetListRes, err error) {
	// 调用service层接口
	out, err := service.Dealer().GetList(ctx, model.DealerGetListInput{
		Page:     req.CommonPaginationReq.Page,
		PageSize: req.CommonPaginationReq.Size,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	res = &apiDealer.DealerGetListRes{
		CommonPaginationRes: apiPagination.CommonPaginationRes{
			Page:  out.Page,
			Size:  out.PageSize,
			Total: out.Total,
			List:  out.Items,
		},
	}
	return
}

// 通过Id查询经销商
func (c *cDealer) DealerGetById(ctx context.Context, req *apiDealer.DealerGetByIdReq) (res *apiDealer.DealerGetByIdRes, err error) {
	// 调用service层接口
	out, err := service.Dealer().GetById(ctx, model.DealerGetByIdInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res = &apiDealer.DealerGetByIdRes{
		DealerInfo: out.DealerInfoBase,
	}
	return
}

// 添加经销商
func (c *cDealer) DealerAdd(ctx context.Context, req *apiDealer.DealerAddReq) (res *apiDealer.DealerAddRes, err error) {
	// 调用service层接口
	out, err := service.Dealer().Add(ctx, model.DealerAddInput{
		DealerAddUpdateBase: model.DealerAddUpdateBase{
			Name:         req.Name,
			Username:     req.Username,
			Phone:        req.Phone,
			DescribeInfo: req.DescribeInfo,
			Province:     req.Province,
			City:         req.City,
			District:     req.District,
			Detail:       req.DetailAddress,
		},
	})
	if err != nil {
		return nil, err
	}
	res = &apiDealer.DealerAddRes{
		Id: out.Id,
	}
	return
}

// 经销商修改
func (c *cDealer) DealerUpdate(ctx context.Context, req *apiDealer.DealerUpdateReq) (res *apiDealer.DealerUpdateRes, err error) {
	// 调用service层接口
	err = service.Dealer().Update(ctx, model.DealerUpdateInput{
		Id: req.Id,
		DealerAddUpdateBase: model.DealerAddUpdateBase{
			Name:         req.Name,
			Username:     req.Username,
			Phone:        req.Phone,
			DescribeInfo: req.DescribeInfo,
			Province:     req.Province,
			City:         req.City,
			District:     req.District,
			Detail:       req.DetailAddress,
		},
	})
	if err != nil {
		return nil, err
	}
	return
}

// 经销商删除
func (c *cDealer) DealerDelete(ctx context.Context, req *apiDealer.DealerDeleteReq) (res *apiDealer.DealerDeleteRes, err error) {
	// 调用service层接口
	err = service.Dealer().Delete(ctx, model.DealerDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}

// 经销商状态修改
func (c *cDealer) DealerUpdateStatus(ctx context.Context, req *apiDealer.DealerUpdateStatusReq) (res *apiDealer.DealerUpdateStatusRes, err error) {
	// 调用service层接口
	err = service.Dealer().UpdateStatus(ctx, model.DealerUpdateStatusInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return
}

// 经销商密码修改
func (c *cDealer) DealerUpdatePassword(ctx context.Context, req *apiDealer.DealerUpdatePasswordReq) (res *apiDealer.DealerUpdatePasswordRes, err error) {
	//调用service接口
	err = service.Dealer().UpdatePassword(ctx, model.DealerUpdatePasswordInput{

		Id:          req.Id,
		OldPassword: req.Password,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return nil, err
	}

	return
}

// 经销商密码重置
func (c *cDealer) DealerResetPassword(ctx context.Context, req *apiDealer.DealerResetPasswordReq) (res *apiDealer.DealerResetPasswordRes, err error) {
	err = service.Dealer().ResetPassword(ctx, model.DealerResetPasswordInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return
}

// 经销商头像修改
func (c *cDealer) DealerUpdateAvatar(ctx context.Context, req *apiDealer.DealerUpdateAvatarReq) (res *apiDealer.DealerUpdateAvatarRes, err error) {
	err = service.Dealer().UpdateAvatar(ctx, model.DealerUpdateAvatarInput{
		Id:  req.Id,
		Url: req.Url,
	})
	if err != nil {
		return nil, err
	}

	return
}
