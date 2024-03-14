package order

import (
	apiOrder "SheeDrive/api/order"
	apiPagination "SheeDrive/api/pagination"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"context"
)

var OrderController = &cOrder{}

type cOrder struct{}

// 订单分页关键字查询
func (c *cOrder) OrderGetList(ctx context.Context, req *apiOrder.OrderGetListReq) (res *apiOrder.OrderGetListRes, err error) {
	out, err := service.Order().GetList(ctx, model.OrderGetListInput{
		Page:       req.CommonPaginationReq.Page,
		PageSize:   req.CommonPaginationReq.Size,
		UserName:   req.UserName,
		DealerName: req.DealerName,
		CarName:    req.CarName,
		Status:     req.Status,
		OrderDate:  req.OrderDate,
	})
	if err != nil {
		return nil, err
	}
	res = &apiOrder.OrderGetListRes{
		CommonPaginationRes: apiPagination.CommonPaginationRes{
			Page:  out.Page,
			Size:  out.PageSize,
			Total: out.Total,
			List:  out.Items,
		},
	}
	return
}

// 订单通过ID查询
func (c *cOrder) OrderGetById(ctx context.Context, req *apiOrder.OrderGetByIdReq) (res *apiOrder.OrderGetByIdRes, err error) {
	out, err := service.Order().GetById(ctx, model.OrderGetByIdInput{Id: req.Id})
	if err != nil {
		return nil, err
	}
	res = &apiOrder.OrderGetByIdRes{
		OrderInfo: out.OrderInfoBase,
	}

	return
}

// 订单添加
func (c *cOrder) OrderAdd(ctx context.Context, req *apiOrder.OrderAddReq) (res *apiOrder.OrderAddRes, err error) {
	out, err := service.Order().Add(ctx, model.OrderAddInput{
		UserId:    req.UserId,
		DealerId:  req.DealerId,
		CarId:     req.CarId,
		AddrId:    req.AddrId,
		OrderTime: req.OrderTime,
	})
	if err != nil {
		return nil, err
	}
	res = &apiOrder.OrderAddRes{
		OrderInfo: out.OrderInfoBase,
	}

	return
}

// 订单删除
func (c *cOrder) OrderDelete(ctx context.Context, req *apiOrder.OrderDeleteReq) (res *apiOrder.OrderDeleteRes, err error) {
	err = service.Order().Delete(ctx, model.OrderDeleteInput{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}

// 订单改为取消状态
func (c *cOrder) OrderUpdateCancel(ctx context.Context, req *apiOrder.OrderUpdateCancelReq) (res *apiOrder.OrderUpdateCancelRes, err error) {
	err = service.Order().UpdateCancel(ctx, model.OrderUpdateInput{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}

// 订单改为确认状态
func (c *cOrder) OrderUpdateConfirm(ctx context.Context, req *apiOrder.OrderUpdateConfirmReq) (res *apiOrder.OrderUpdateConfirmRes, err error) {
	err = service.Order().UpdateConfirm(ctx, model.OrderUpdateInput{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}

// 订单改为签订协议状态
func (c *cOrder) OrderUpdateSign(ctx context.Context, req *apiOrder.OrderUpdateSignReq) (res *apiOrder.OrderUpdateSignRes, err error) {
	err = service.Order().UpdateSign(ctx, model.OrderUpdateInput{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}

// 订单改为开始状态
func (c *cOrder) OrderUpdateStart(ctx context.Context, req *apiOrder.OrderUpdateStartReq) (res *apiOrder.OrderUpdateStartRes, err error) {
	err = service.Order().UpdateStart(ctx, model.OrderUpdateInput{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}

// 订单改为结束状态
func (c *cOrder) OrderUpdateEnd(ctx context.Context, req *apiOrder.OrderUpdateEndReq) (res *apiOrder.OrderUpdateEndRes, err error) {
	err = service.Order().UpdateEnd(ctx, model.OrderUpdateInput{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}

// 订单改为待评价状态
func (c *cOrder) OrderUpdateEndAll(ctx context.Context, req *apiOrder.OrderUpdateEndAllReq) (res *apiOrder.OrderUpdateEndAllRes, err error) {
	err = service.Order().UpdateEndAll(ctx, model.OrderUpdateInput{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}

// 获取汽车排行
func (c *cOrder) OrderGetCarRank(ctx context.Context, req *apiOrder.OrderGetCarRankReq) (res *apiOrder.OrderGetCarRankRes, err error) {
	out, err := service.Order().GetCarRank(ctx)
	if err != nil {
		return nil, err
	}
	res = &apiOrder.OrderGetCarRankRes{
		CommonPaginationRes: apiPagination.CommonPaginationRes{
			Page:  1,
			Size:  10,
			Total: 10,
			List:  out.Items,
		},
	}
	return
}

func (c *cOrder) OrderGetIncomplete(ctx context.Context, req *apiOrder.OrderGetIncompleteReq) (res *apiOrder.OrderGetIncompleteRes, err error) {
	out, err := service.Order().GetIncomplete(ctx, model.OrderGetIncompleteInput{
		DealerId: req.DealerId,
	})
	if err != nil {
		return nil, err
	}
	res = &apiOrder.OrderGetIncompleteRes{
		Total: out.Total,
	}
	return
}

// 获取地址出现次数
func (c *cOrder) OrderGetAddressTimes(ctx context.Context, req *apiOrder.OrderGetAddressTimesReq) (res *apiOrder.OrderGetAddressTimesRes, err error) {
	out, err := service.Order().GetAddressTimes(ctx, model.OrderGetAddressTimesInput{DealerId: req.DealerId})
	if err != nil {
		return nil, err
	}
	res = &apiOrder.OrderGetAddressTimesRes{
		CommonPaginationRes: apiPagination.CommonPaginationRes{
			Page:  1,
			Size:  10,
			Total: 10,
			List:  out.Items,
		},
	}
	return
}

func (c *cOrder) OrderGetTimeCount(ctx context.Context, req *apiOrder.OrderGetTimeCountReq) (res *apiOrder.OrderGetTimeCountRes, err error) {
	out, err := service.Order().GetTimeCount(ctx, model.OrderGetTimeCountInput{DealerId: req.DealerId})
	if err != nil {
		return nil, err
	}
	res = &apiOrder.OrderGetTimeCountRes{
		TimeSeries: out.TimeSeries,
		OrderCount: out.OrderCount,
	}

	return
}

func (c *cOrder) OrderGetByUserId(ctx context.Context, req *apiOrder.OrderGetByUserIdReq) (res *apiOrder.OrderGetByUserIdRes, err error) {
	out, err := service.Order().GetByUserId(ctx, model.OrderGetByUserIdInput{
		Page:     req.CommonPaginationReq.Page,
		PageSize: req.CommonPaginationReq.Size,
		UserId:   req.UserId,
	})
	if err != nil {
		return nil, err
	}
	res = &apiOrder.OrderGetByUserIdRes{
		CommonPaginationRes: apiPagination.CommonPaginationRes{
			Page:  out.Page,
			Size:  out.PageSize,
			Total: out.Total,
			List:  out.Items,
		},
	}
	return
}
