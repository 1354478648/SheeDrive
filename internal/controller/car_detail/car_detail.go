package cardetail

import (
	apiCarDetail "SheeDrive/api/car_detail"
	apiPagination "SheeDrive/api/pagination"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"context"
)

var CarDetailController = &cCarDetail{}

type cCarDetail struct{}

// 汽车信息分页关键字查询
func (c *cCarDetail) CarDetailGetListReq(ctx context.Context, req *apiCarDetail.CarDetailGetListReq) (res *apiCarDetail.CarDetailGetListRes, err error) {
	out, err := service.CarDetail().GetList(ctx, model.CarDetailGetListInput{
		Page:      req.CommonPaginationReq.Page,
		PageSize:  req.CommonPaginationReq.Size,
		Year:      req.Year,
		Brand:     req.Brand,
		Model:     req.Model,
		Category:  req.Category,
		LowPrice:  req.LowPrice,
		HighPrice: req.HighPrice,
	})
	if err != nil {
		return nil, err
	}
	res = &apiCarDetail.CarDetailGetListRes{
		CommonPaginationRes: apiPagination.CommonPaginationRes{
			Page:  out.Page,
			Size:  out.PageSize,
			Total: out.Total,
			List:  out.Items,
		},
	}
	return
}
