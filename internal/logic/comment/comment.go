package comment

import (
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/model/do"
	"SheeDrive/internal/service"
	"SheeDrive/utility"
	"context"

	"github.com/gogf/gf/errors/gerror"
)

type iComment struct{}

func New() *iComment {
	return &iComment{}
}

func init() {
	service.RegisterComment(New())
}

// GetList implements service.IComment.
func (i *iComment) GetList(ctx context.Context, in model.CommentGetListInput) (out *model.CommentGetListOutput, err error) {
	out = &model.CommentGetListOutput{
		Page:     in.Page,
		PageSize: in.PageSize,
	}

	var (
		md = dao.Comment.Ctx(ctx)
	)

	if in.OrderId != 0 {
		md = md.Where(dao.Comment.Columns().OrderId, in.OrderId)
	}
	if in.DealerName != "" {
		dealerId, err := dao.Dealer.Ctx(ctx).Fields("id").WhereLike(dao.Dealer.Columns().Name, "%"+in.DealerName+"%").Array()
		if err != nil {
			return out, err
		}
		orderId, err := dao.Order.Ctx(ctx).Fields("id").WhereIn(dao.Order.Columns().DealerId, dealerId).Array()
		if err != nil {
			return out, err
		}
		md = md.WhereIn(dao.Comment.Columns().OrderId, orderId)
	}
	if in.CarName != "" {
		carId, err := dao.CarDetail.Ctx(ctx).Fields("id").WhereLike("CONCAT(year, brand, model, version)", "%"+in.CarName+"%").Array()
		if err != nil {
			return out, err
		}
		orderId, err := dao.Order.Ctx(ctx).Fields("id").WhereIn(dao.Order.Columns().CarId, carId).Array()
		if err != nil {
			return out, err
		}
		md = md.WhereIn(dao.Comment.Columns().OrderId, orderId)
	}
	if (in.BeforeDate != nil) && (in.AfterDate != nil) {
		md = md.WhereBetween(dao.Comment.Columns().CreateTime, in.BeforeDate, in.AfterDate)
	}

	md = md.OrderDesc(dao.Comment.Columns().CreateTime).Page(in.Page, in.PageSize)

	out.Total, err = md.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	md.WithAll().Scan(&out.Items)

	return
}

// GetById implements service.IComment.
func (i *iComment) GetById(ctx context.Context, in model.CommentGetByIdInput) (out *model.CommentGetByIdOutput, err error) {
	out = &model.CommentGetByIdOutput{}

	err = dao.Comment.Ctx(ctx).WithAll().Where(dao.Comment.Columns().Id, in.Id).Scan(&out.CommentInfoBase)
	if err != nil {
		return nil, gerror.New("该评论信息不存在")
	}
	return
}

// Add implements service.IComment.
func (i *iComment) Add(ctx context.Context, in model.CommentAddInput) (out *model.CommentAddOutput, err error) {
	out = &model.CommentAddOutput{}

	id := utility.GenSnowFlakeId()
	_, err = dao.Comment.Ctx(ctx).Data(do.Comment{
		Id:          id,
		OrderId:     in.OrderId,
		Content:     in.Content,
		TotalScore:  in.TotalScore,
		DealerScore: in.DealerScore,
		CarScore:    in.CarScore,
	}).Insert()
	if err != nil {
		return nil, gerror.New("同一订单不允许被评价多次")
	}
	out.Id = id

	return
}

// Delete implements service.IComment.
func (i *iComment) Delete(ctx context.Context, in model.CommentDeleteInput) (err error) {
	_, err = dao.Comment.Ctx(ctx).Where(dao.Comment.Columns().Id, in.Id).Delete()
	if err != nil {
		return gerror.New("删除评价失败")
	}
	return
}

// GetAvg implements service.IComment.
func (i *iComment) GetAvg(ctx context.Context, in model.CommentGetAvgInput) (out *model.CommentGetAvgOutput, err error) {
	out = &model.CommentGetAvgOutput{}

	orderId, err := dao.Order.Ctx(ctx).Fields(dao.Order.Columns().Id).Where(dao.Order.Columns().DealerId, in.DealerId).Array()
	if err != nil {
		return nil, gerror.New("获取订单列表失败")
	}
	avg, err := dao.Comment.Ctx(ctx).WhereIn(dao.Comment.Columns().OrderId, orderId).Avg(dao.Comment.Columns().DealerScore)

	out.Avg = avg

	return
}
