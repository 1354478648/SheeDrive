package order

import (
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"context"
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
