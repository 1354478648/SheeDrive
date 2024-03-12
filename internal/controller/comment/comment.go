package comment

import (
	apiComment "SheeDrive/api/comment"
	apiPagination "SheeDrive/api/pagination"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"context"
)

var CommentController = &cComment{}

type cComment struct{}

// 经销商列表查询
func (c *cComment) CommentGetList(ctx context.Context, req *apiComment.CommentGetListReq) (res *apiComment.CommentGetListRes, err error) {
	// 调用service层接口
	out, err := service.Comment().GetList(ctx, model.CommentGetListInput{
		Page:       req.CommonPaginationReq.Page,
		PageSize:   req.CommonPaginationReq.Size,
		OrderId:    req.OrderId,
		DealerName: req.DealerName,
		CarName:    req.CarName,
		BeforeDate: req.BeforeDate,
		AfterDate:  req.AfterDate,
	})
	if err != nil {
		return nil, err
	}
	res = &apiComment.CommentGetListRes{
		CommonPaginationRes: apiPagination.CommonPaginationRes{
			Page:  out.Page,
			Size:  out.PageSize,
			Total: out.Total,
			List:  out.Items,
		},
	}
	return
}

func (c *cComment) CommentGetById(ctx context.Context, req *apiComment.CommentGetByIdReq) (res *apiComment.CommentGetByIdRes, err error) {
	out, err := service.Comment().GetById(ctx, model.CommentGetByIdInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res = &apiComment.CommentGetByIdRes{
		CommentInfo: out.CommentInfoBase,
	}
	return
}

func (c *cComment) CommentAdd(ctx context.Context, req *apiComment.CommentAddReq) (res *apiComment.CommentAddRes, err error) {
	out, err := service.Comment().Add(ctx, model.CommentAddInput{
		OrderId:     req.OrderId,
		Content:     req.Content,
		TotalScore:  req.TotalScore,
		DealerScore: req.DealerScore,
		CarScore:    req.CarScore,
	})
	if err != nil {
		return nil, err
	}
	res = &apiComment.CommentAddRes{
		Id: out.Id,
	}
	return
}

func (c *cComment) CommentDelete(ctx context.Context, req *apiComment.CommentDeleteReq) (res *apiComment.CommentDeleteRes, err error) {
	err = service.Comment().Delete(ctx, model.CommentDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (c *cComment) CommentGetAvg(ctx context.Context, req *apiComment.CommentGetAvgReq) (res *apiComment.CommentGetAvgRes, err error) {
	out, err := service.Comment().GetAvg(ctx, model.CommentGetAvgInput{
		DealerId: req.DealerId,
	})
	if err != nil {
		return nil, err
	}
	res = &apiComment.CommentGetAvgRes{
		Avg: out.Avg,
	}

	return
}
