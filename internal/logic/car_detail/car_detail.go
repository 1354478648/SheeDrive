package cardetail

import (
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"context"
)

type iCarDetail struct{}

func New() *iCarDetail {
	return &iCarDetail{}
}

func init() {
	service.RegisterCarDetail(New())
}

// GetList implements service.ICarDetail.
func (*iCarDetail) GetList(ctx context.Context, in model.CarDetailGetListInput) (out *model.CarDetailGetListOutput, err error) {
	// 1. 实例化响应结构体
	out = &model.CarDetailGetListOutput{
		Page:     in.Page,
		PageSize: in.PageSize,
	}
	// 2.获取*gdb.Model对象
	var (
		md = dao.CarDetail.Ctx(ctx)
	)

	// 3.通过md构造动态SQL语句
	// 3.1 判断是否有关键字Year查询
	if in.Year != "" {
		md = md.Where(dao.CarDetail.Columns().Year, in.Year)
	}
	// 3.2 判断是否有关键字Brand查询
	if in.Brand != "" {
		md = md.Where(dao.CarDetail.Columns().Brand, in.Brand)
	}
	// 3.3 判断是否有关键字Model查询
	if in.Model != "" {
		md = md.WhereLike(dao.CarDetail.Columns().Model, "%"+in.Model+"%")
	}
	// 3.4 判断是否有关键字Category查询
	if in.Category != "" {
		md = md.Where(dao.CarDetail.Columns().Category, in.Category)
	}
	// 3.5 判断是否有关键字LowPrice和HighPrice查询
	if in.LowPrice != nil || in.HighPrice != nil {
		md = md.WhereBetween(dao.CarDetail.Columns().Price, in.LowPrice, in.HighPrice)
	}

	// 4. 执行分页查询
	// 设置排序：更新时间降序
	md = md.OrderDesc(dao.CarDetail.Columns().UpdateTime)
	// 设置分页
	md = md.Page(in.Page, in.PageSize)

	// 5. 判断当前页的数据条数
	out.Total, err = md.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	// 6. 将查询结果赋值给响应结构体
	if err := md.Scan(&out.Items); err != nil {
		return out, err
	}

	return
}
