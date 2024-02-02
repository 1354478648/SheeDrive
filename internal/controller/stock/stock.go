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
