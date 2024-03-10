package cardetail

import (
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/model/do"
	"SheeDrive/internal/service"
	"context"

	"github.com/gogf/gf/errors/gerror"
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
	// 3.5 判断是否有关键字Type查询
	if in.Type != "" {
		md = md.Where(dao.CarDetail.Columns().Type, in.Type)
	}
	// 3.6 判断是否有关键字LowPrice和HighPrice查询
	if in.LowPrice != 0 && in.HighPrice != 0 {
		md = md.WhereBetween(dao.CarDetail.Columns().Price, in.LowPrice, in.HighPrice)
	}
	if in.LowPrice != 0 {
		md = md.WhereGTE(dao.CarDetail.Columns().Price, in.LowPrice)
	}
	if in.HighPrice != 0 {
		md = md.WhereLTE(dao.CarDetail.Columns().Price, in.HighPrice)
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

// GetById implements service.ICarDetail.
func (*iCarDetail) GetById(ctx context.Context, in model.CarDetailGetByIdInput) (out *model.CarDetailGetByIdOutput, err error) {
	// 实例化响应结构体
	out = &model.CarDetailGetByIdOutput{}

	err = dao.CarDetail.Ctx(ctx).Where(dao.CarDetail.Columns().Id, in.Id).Scan(&out.CarDetail)
	if err != nil {
		return out, gerror.New("该汽车信息不存在")
	}
	return
}

// Add implements service.ICarDetail.
func (*iCarDetail) Add(ctx context.Context, in model.CarDetailAddInput) (out *model.CarDetailAddOutput, err error) {
	// 实例化响应结构体
	out = &model.CarDetailAddOutput{}

	// 执行添加操作
	id, err := dao.CarDetail.Ctx(ctx).Data(do.CarDetail{
		Year:         in.Year,
		Brand:        in.Brand,
		Model:        in.Model,
		Version:      in.Version,
		Image:        in.Image,
		Category:     in.Category,
		Price:        in.Price,
		Type:         in.Type,
		Seats:        in.Seats,
		DescribeInfo: in.DescribeInfo,
	}).InsertAndGetId()
	if err != nil {
		return out, gerror.New("汽车信息添加失败")
	}

	// 将自增主键id赋值给响应结构体
	out.Id = id

	return
}

// Update implements service.ICarDetail.
func (*iCarDetail) Update(ctx context.Context, in model.CarDetailUpdateInput) (err error) {
	_, err = dao.CarDetail.Ctx(ctx).Where(dao.CarDetail.Columns().Id, in.Id).Update(do.CarDetail{
		Year:         in.Year,
		Brand:        in.Brand,
		Model:        in.Model,
		Version:      in.Version,
		Image:        in.Image,
		Category:     in.Category,
		Price:        in.Price,
		Type:         in.Type,
		Seats:        in.Seats,
		DescribeInfo: in.DescribeInfo,
	})
	if err != nil {
		return gerror.New("汽车信息更新失败")
	}

	return
}

// Delete implements service.ICarDetail.
func (*iCarDetail) Delete(ctx context.Context, in model.CarDetailDeleteInput) (err error) {
	_, err = dao.CarDetail.Ctx(ctx).Where(dao.CarDetail.Columns().Id, in.Id).Delete()
	if err != nil {
		return gerror.New("汽车信息删除失败")
	}

	// 执行删除库存操作
	_, err = dao.Stock.Ctx(ctx).Where(dao.Stock.Columns().CarId, in.Id).Delete()
	if err != nil {
		return gerror.New("删除汽车库存失败")
	}

	// 执行删除评价操作
	orderId, err := dao.Order.Ctx(ctx).Fields("id").Where(dao.Order.Columns().CarId, in.Id).Array()
	if err != nil {
		return gerror.New("未找到该汽车下的订单")
	}
	_, err = dao.Comment.Ctx(ctx).WhereIn(dao.Comment.Columns().OrderId, orderId).Delete()
	if err != nil {
		return gerror.New("删除评价失败")
	}

	// 执行删除订单操作
	_, err = dao.Order.Ctx(ctx).Where(dao.Order.Columns().CarId, in.Id).Delete()
	if err != nil {
		return gerror.New("删除订单失败")
	}

	return
}
