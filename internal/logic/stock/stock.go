package stock

import (
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/model/do"
	"SheeDrive/internal/service"
	"context"

	"github.com/gogf/gf/errors/gerror"
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
	if in.CarName != "" {
		carId, err := dao.CarDetail.Ctx(ctx).Fields("id").WhereLike("CONCAT(year, brand, model, version)", "%"+in.CarName+"%").Array()
		if err != nil {
			return out, err
		}
		md = md.WhereIn(dao.Stock.Columns().CarId, carId)
	}

	// 4. 设置排序和分页
	md = md.OrderDesc(dao.Stock.Columns().CreateTime).Page(in.Page, in.PageSize)

	// 5. 计算数据条数
	out.Total, err = md.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	// 6. 关联查询
	md.WithAll().Scan(&out.Items)

	return
}

// GetById implements service.IStock.
func (*iStock) GetById(ctx context.Context, in model.StockGetByIdInput) (out *model.StockGetByIdOutput, err error) {
	out = &model.StockGetByIdOutput{}

	err = dao.Stock.Ctx(ctx).WithAll().Where(dao.Stock.Columns().Id, in.Id).Scan(&out.StockInfoBase)
	if err != nil {
		return nil, gerror.New("该库存信息不存在")
	}
	return
}

// Add implements service.IStock.
func (*iStock) Add(ctx context.Context, in model.StockAddInput) (out *model.StockAddOutput, err error) {
	out = &model.StockAddOutput{}

	id, err := dao.Stock.Ctx(ctx).Data(do.Stock{
		DealerId: in.DealerId,
		CarId:    in.CarId,
	}).InsertAndGetId()
	if err != nil {
		// 经销商和汽车信息字段组合需要有唯一性索引
		// ALTER TABLE `stock` ADD CONSTRAINT `uc_dealer_car` UNIQUE (`dealer_id`, `car_id`)
		return out, gerror.New("不允许重复添加已存在的库存信息")
	}
	out.Id = id

	return
}

// Delete implements service.IStock.
func (*iStock) Delete(ctx context.Context, in model.StockDeleteInput) (err error) {
	_, err = dao.Stock.Ctx(ctx).Where(dao.Stock.Columns().Id, in.Id).Delete()
	if err != nil {
		return gerror.New("删除库存失败")
	}
	return
}
