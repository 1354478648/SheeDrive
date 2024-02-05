package stock

import (
	apiPagination "SheeDrive/api/pagination"
	apiStock "SheeDrive/api/stock"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"context"
)

var StockController = &cStock{}

type cStock struct{}

// 库存分页关键字查询（管理员）
func (c *cStock) GetList(ctx context.Context, req *apiStock.StockGetListReq) (res *apiStock.StockGetListRes, err error) {
	out, err := service.Stock().GetList(ctx, model.StockGetListInput{
		Page:       req.CommonPaginationReq.Page,
		PageSize:   req.CommonPaginationReq.Size,
		DealerName: req.DealerName,
	})
	if err != nil {
		return nil, err
	}
	res = &apiStock.StockGetListRes{
		CommonPaginationRes: apiPagination.CommonPaginationRes{
			Page:  out.Page,
			Size:  out.PageSize,
			Total: out.Total,
			List:  out.Items,
		},
	}
	return
}

// 库存通过Id查询
func (c *cStock) StockGetById(ctx context.Context, req *apiStock.StockGetByIdReq) (res *apiStock.StockGetByIdRes, err error) {
	out, err := service.Stock().GetById(ctx, model.StockGetByIdInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res = &apiStock.StockGetByIdRes{
		StockInfo: out.StockInfoBase,
	}

	return
}

// 库存添加
func (c *cStock) StockAdd(ctx context.Context, req *apiStock.StockAddReq) (res *apiStock.StockAddRes, err error) {
	out, err := service.Stock().Add(ctx, model.StockAddInput{
		DealerId: req.DealerId,
		CarId:    req.CarId,
	})
	if err != nil {
		return nil, err
	}
	res = &apiStock.StockAddRes{
		Id: out.Id,
	}
	return
}

// 库存更新只有一个字段需要更新，没有实现的意义

// 库存删除
func (c *cStock) StockDelete(ctx context.Context, req *apiStock.StockDeleteReq) (res *apiStock.StockDeleteRes, err error) {
	err = service.Stock().Delete(ctx, model.StockDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}
