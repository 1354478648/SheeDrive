package stock

import (
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"context"
)

type iStock struct{}

func New() *iStock {
	return &iStock{}
}

func init() {
	service.RegisterStock(New())
}

// GetList implements service.IStock.
func (*iStock) GetList(ctx context.Context, in model.StockGetListInput) (out *model.StockGetListOutput, err error) {
	// 1. 实例化响应结构体
	out = &model.StockGetListOutput{
		Page:     in.Page,
		PageSize: in.PageSize,
	}

	// 2. 获取*gdb.Model对象
	var (
		md = dao.Stock.Ctx(ctx)
	)

	// 3. 构造动态SQL语句
	if in.DealerName != "" {
		dealerId, err := dao.Dealer.Ctx(ctx).Fields("id").WhereLike(dao.Dealer.Columns().Name, "%"+in.DealerName+"%").Array()
		if err != nil {
			return out, err
		}
		md = md.WhereIn(dao.Stock.Columns().DealerId, dealerId)
	}

	// 4. 设置排序和分页
	md = md.OrderDesc(dao.Stock.Columns().UpdateTime).Page(in.Page, in.PageSize)

	// 5. 计算数据条数
	out.Total, err = md.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	// 6. 关联查询
	md.WithAll().Scan(&out.Items)

	return
}
