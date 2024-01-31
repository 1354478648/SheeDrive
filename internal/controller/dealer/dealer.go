package dealer

import (
	apiDealer "SheeDrive/api/dealer"
	apiPagination "SheeDrive/api/pagination"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"SheeDrive/utility"
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
		Token:      utility.GenToken(dealer.Username),
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
		Name:          req.Name,
		Username:      req.Username,
		Phone:         req.Phone,
		DescribeInfo:  req.DescribeInfo,
		Province:      req.Province,
		City:          req.City,
		District:      req.District,
		DetailAddress: req.DetailAddress,
	})
	if err != nil {
		return nil, err
	}
	res = &apiDealer.DealerAddRes{
		Id: out.Id,
	}
	return
}
