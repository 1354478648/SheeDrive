package order

import (
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/model/do"
	"SheeDrive/internal/service"
	"SheeDrive/utility"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type iOrder struct{}

func New() *iOrder {
	return &iOrder{}
}

func init() {
	service.RegisterOrder(New())
}

// GetList implements service.IOrder.
func (i *iOrder) GetList(ctx context.Context, in model.OrderGetListInput) (out *model.OrderGetListOutput, err error) {
	// 1. 实例化响应结构体
	out = &model.OrderGetListOutput{
		Page:     in.Page,
		PageSize: in.PageSize,
	}

	// 2. 获取*gdb.Model对象
	var (
		md = dao.Order.Ctx(ctx)
	)

	// 3. 构造动态SQL语句
	if in.UserName != "" {
		userId, err := dao.User.Ctx(ctx).Fields("id").WhereLike("CONCAT(last_name, first_name)", "%"+in.UserName+"%").Array()
		if err != nil {
			return out, err
		}
		md = md.WhereIn(dao.Order.Columns().UserId, userId)
	}
	if in.DealerName != "" {
		dealerId, err := dao.Dealer.Ctx(ctx).Fields("id").WhereLike(dao.Dealer.Columns().Name, "%"+in.DealerName+"%").Array()
		if err != nil {
			return out, err
		}
		md = md.WhereIn(dao.Order.Columns().DealerId, dealerId)
	}
	if in.CarName != "" {
		carId, err := dao.CarDetail.Ctx(ctx).Fields("id").WhereLike("CONCAT(year, brand, model, version)", "%"+in.CarName+"%").Array()
		if err != nil {
			return out, err
		}
		md = md.WhereIn(dao.Order.Columns().CarId, carId)
	}
	if in.Status != -2 {
		md = md.WhereIn(dao.Order.Columns().Status, in.Status)
	}
	if in.OrderDate != nil {
		md = md.WhereIn(dao.Order.Columns().OrderTime, in.OrderDate)
	}

	// 4. 设置排序和分页
	md = md.OrderDesc(dao.Order.Columns().CreateTime).Page(in.Page, in.PageSize)

	// 5. 计算数据条数
	out.Total, err = md.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	// 6. 关联查询
	md.WithAll().Scan(&out.Items)

	return
}

// GetById implements service.IOrder.
func (i *iOrder) GetById(ctx context.Context, in model.OrderGetByIdInput) (out *model.OrderGetByIdOutput, err error) {
	out = &model.OrderGetByIdOutput{}

	err = dao.Order.Ctx(ctx).WithAll().Where(dao.Order.Columns().Id, in.Id).Scan(&out.OrderInfoBase)
	if err != nil {
		return nil, gerror.New("该订单信息不存在")
	}

	return
}

// Add implements service.IOrder.
func (i *iOrder) Add(ctx context.Context, in model.OrderAddInput) (out *model.OrderAddOutput, err error) {
	out = &model.OrderAddOutput{}

	id := utility.GenSnowFlakeId()
	_, err = dao.Order.Ctx(ctx).Data(do.Order{
		Id:        id,
		UserId:    in.UserId,
		DealerId:  in.DealerId,
		CarId:     in.CarId,
		AddrId:    in.AddrId,
		Status:    1,
		OrderTime: in.OrderTime,
	}).Insert()
	if err != nil {
		return out, gerror.New("订单生成失败")
	}
	order, err := service.Order().GetById(ctx, model.OrderGetByIdInput{Id: id})
	if err != nil {
		return out, gerror.New("未找到该订单")
	}
	out.OrderInfoBase = order.OrderInfoBase

	return
}

// Delete implements service.IOrder.
func (i *iOrder) Delete(ctx context.Context, in model.OrderDeleteInput) (err error) {
	// 执行删除评价操作
	_, err = dao.Comment.Ctx(ctx).WhereIn(dao.Comment.Columns().OrderId, in.Id).Delete()
	if err != nil {
		return gerror.New("删除评价失败")
	}

	_, err = dao.Order.Ctx(ctx).Where(dao.Order.Columns().Id, in.Id).Delete()
	if err != nil {
		return gerror.New("删除订单失败")
	}
	return
}

// UpdateCancel implements service.IOrder.
func (i *iOrder) UpdateCancel(ctx context.Context, in model.OrderUpdateInput) (err error) {
	_, err = dao.Order.Ctx(ctx).Where(dao.Order.Columns().Id, in.Id).Data(do.Order{
		Status: 0,
	}).Update()
	if err != nil {
		return gerror.New("订单取消失败")
	}
	return
}

// UpdateConfirm implements service.IOrder.
func (i *iOrder) UpdateConfirm(ctx context.Context, in model.OrderUpdateInput) (err error) {
	_, err = dao.Order.Ctx(ctx).Where(dao.Order.Columns().Id, in.Id).Data(do.Order{
		Status:      2,
		ConfirmTime: gtime.Now(),
	}).Update()
	if err != nil {
		return gerror.New("订单确认失败")
	}
	return
}

// UpdateEnd implements service.IOrder.
func (i *iOrder) UpdateEnd(ctx context.Context, in model.OrderUpdateInput) (err error) {
	_, err = dao.Order.Ctx(ctx).Where(dao.Order.Columns().Id, in.Id).Data(do.Order{
		Status:  5,
		EndTime: gtime.Now(),
	}).Update()
	if err != nil {
		return gerror.New("订单结束试驾失败")
	}
	return
}

// UpdateSign implements service.IOrder.
func (i *iOrder) UpdateSign(ctx context.Context, in model.OrderUpdateInput) (err error) {
	_, err = dao.Order.Ctx(ctx).Where(dao.Order.Columns().Id, in.Id).Data(do.Order{
		Status:   3,
		SignTime: gtime.Now(),
	}).Update()
	if err != nil {
		return gerror.New("订单签署协议失败")
	}
	return
}

// UpdateStart implements service.IOrder.
func (i *iOrder) UpdateStart(ctx context.Context, in model.OrderUpdateInput) (err error) {
	_, err = dao.Order.Ctx(ctx).Where(dao.Order.Columns().Id, in.Id).Data(do.Order{
		Status:    4,
		StartTime: gtime.Now(),
	}).Update()
	if err != nil {
		return gerror.New("订单开始试驾失败")
	}
	return
}

// UpdateEndAll implements service.IOrder.
func (i *iOrder) UpdateEndAll(ctx context.Context, in model.OrderUpdateInput) (err error) {
	_, err = dao.Order.Ctx(ctx).Where(dao.Order.Columns().Id, in.Id).Data(do.Order{
		Status:         6,
		PrecommentTime: gtime.Now(),
	}).Update()
	if err != nil {
		return gerror.New("订单切换至待评价状态失败")
	}
	return
}
