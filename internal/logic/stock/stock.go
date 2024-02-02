package stock

import (
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/model/do"
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
	// 3. 执行关联查询
	// md = md.With(model.DealerInfoBase{}, do.CarDetail{})
	dealer_md := dao.Dealer.Ctx(ctx)
	var stock model.StockInfoBase

	md = md.With(do.CarDetail{})

	// 动态SQL
	// if in.DealerName != "" {
	// 	md = md.WhereLike(dao.Dealer.Columns().Name, "%"+in.DealerName+"%")
	// }

	// 4. 执行分页查询
	// 设置排序：更新时间降序
	md = md.OrderDesc(dao.Stock.Columns().UpdateTime)
	// 设置分页
	md = md.Page(in.Page, in.PageSize)

	// 5. 判断当前页的数据条数
	out.Total, err = md.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	_ = dealer_md.ScanList(&stock.Dealer, "id", stock.DealerId)
	// 6. 将查询结果赋值给响应结构体
	if err := md.Scan(&out.Items); err != nil {
		return out, err

	}
	return
}
