package cardetail

import (
	apiCarDetail "SheeDrive/api/car_detail"
	apiPagination "SheeDrive/api/pagination"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

var CarDetailController = &cCarDetail{}

type cCarDetail struct{}

// 汽车信息分页关键字查询
func (c *cCarDetail) CarDetailGetListReq(ctx context.Context, req *apiCarDetail.CarDetailGetListReq) (res *apiCarDetail.CarDetailGetListRes, err error) {
	if req.LowPrice != 0 && req.HighPrice != 0 {
		if req.LowPrice > req.HighPrice {
			return nil, gerror.New("最低价格不能大于最高价格")
		}
	}

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

// 汽车信息通过Id查询
func (c *cCarDetail) CarDetailGetByIdReq(ctx context.Context, req *apiCarDetail.CarDetailGetByIdReq) (res *apiCarDetail.CarDetailGetByIdRes, err error) {
	out, err := service.CarDetail().GetById(ctx, model.CarDetailGetByIdInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res = &apiCarDetail.CarDetailGetByIdRes{
		CarDetail: out.CarDetail,
	}
	return
}

// 汽车信息添加
func (c *cCarDetail) CarDetailAddReq(ctx context.Context, req *apiCarDetail.CarDetailAddReq) (res *apiCarDetail.CarDetailAddRes, err error) {
	out, err := service.CarDetail().Add(ctx, model.CarDetailAddInput{
		CarDetailAddUpdateBase: model.CarDetailAddUpdateBase{
			Year:         req.Year,
			Brand:        req.Brand,
			Model:        req.Model,
			Version:      req.Version,
			Image:        req.Image,
			Category:     req.Category,
			Color:        req.Color,
			Price:        req.Price,
			Type:         req.Type,
			Seats:        req.Seats,
			DescribeInfo: req.DescribeInfo,
		},
	})
	if err != nil {
		return nil, err
	}
	res = &apiCarDetail.CarDetailAddRes{
		Id: out.Id,
	}
	return
}

// 汽车信息修改
func (c *cCarDetail) CarDetailUpdateReq(ctx context.Context, req *apiCarDetail.CarDetailUpdateReq) (res *apiCarDetail.CarDetailUpdateRes, err error) {
	err = service.CarDetail().Update(ctx, model.CarDetailUpdateInput{
		Id: req.Id,
		CarDetailAddUpdateBase: model.CarDetailAddUpdateBase{
			Year:         req.Year,
			Brand:        req.Brand,
			Model:        req.Model,
			Version:      req.Version,
			Image:        req.Image,
			Category:     req.Category,
			Color:        req.Color,
			Price:        req.Price,
			Type:         req.Type,
			Seats:        req.Seats,
			DescribeInfo: req.DescribeInfo,
		},
	},
	)
	if err != nil {
		return nil, err
	}
	return
}

// 汽车信息删除
func (c *cCarDetail) CarDetailDeleteReq(ctx context.Context, req *apiCarDetail.CarDetailDeleteReq) (res *apiCarDetail.CarDetailDeleteRes, err error) {
	err = service.CarDetail().Delete(ctx, model.CarDetailDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}
