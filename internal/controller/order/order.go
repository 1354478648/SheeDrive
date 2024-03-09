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
