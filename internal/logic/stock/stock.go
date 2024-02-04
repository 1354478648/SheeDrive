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
// func (*iStock) GetList(ctx context.Context, in model.StockGetListInput) (out *model.StockGetListOutput, err error) {
// 	// 1. 实例化响应结构体
// 	out = &model.StockGetListOutput{
// 		Page:     in.Page,
// 		PageSize: in.PageSize,
// 	}

// 	// // 2. 获取*gdb.Model对象
// 	var (
// 		md = dao.Stock.Ctx(ctx)
// 	)

// 	// // 动态SQL
// 	// if in.DealerName != "" {
// 	// 	md = md.WhereLike(dao.Dealer.Columns().Name, "%"+in.DealerName+"%")
// 	// }

// 	// 4. 执行分页查询
// 	//设置排序：更新时间降序
// 	md = md.OrderDesc(dao.Stock.Columns().UpdateTime)
// 	// 设置分页
// 	md = md.Page(in.Page, in.PageSize)

// 	// 3. 执行关联查询
// 	var stockList []model.StockInfoBase

// 	err = md.ScanList(&stockList, "StockInfo")
// 	if err != nil {
// 		return nil, err
// 	}

// 	if in.DealerName != "" {
// 		err = dao.Dealer.Ctx(ctx).
// 			Where("id", gdb.ListItemValuesUnique(stockList, "StockInfo", "DealerId")).
// 			WhereLike(dao.Dealer.Columns().Name, "%"+in.DealerName+"%").
// 			ScanList(&stockList, "Dealer", "StockInfo", "id:DealerId")
// 		if err != nil {
// 			return nil, err
// 		}
// 	} else {
// 		err = dao.Dealer.Ctx(ctx).
// 			Where("id", gdb.ListItemValuesUnique(stockList, "StockInfo", "DealerId")).
// 			ScanList(&stockList, "Dealer", "StockInfo", "id:DealerId")
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	err = dao.CarDetail.Ctx(ctx).
// 		Where("id", gdb.ListItemValuesUnique(stockList, "StockInfo", "CarId")).
// 		ScanList(&stockList, "CarDetail", "StockInfo", "id:CarId")
// 	if err != nil {
// 		return nil, err
// 	}

// 	for i := 0; i < len(stockList); {
// 		if stockList[i].Dealer == nil {
// 			// 如果 Dealer 属性为 nil，则从切片中移除该对象
// 			stockList = append(stockList[:i], stockList[i+1:]...)
// 		} else {
// 			// 如果 Dealer 属性不为 nil，则继续下一个对象
// 			i++
// 		}
// 	}
// 	// 6. 将查询结果赋值给响应结构体
// 	out.Items = stockList

// 	// 5. 判断当前页的数据条数
// 	out.Total = len(stockList)
// 	if len(stockList) == 0 {
// 		return out, err
// 	}
// 	return
// }

// GetList implements service.IStock.
func (*iStock) GetList(ctx context.Context, in model.StockGetListInput) (out *model.StockGetListOutput, err error) {
	// 1. 实例化响应结构体
	out = &model.StockGetListOutput{
		Page:     in.Page,
		PageSize: in.PageSize,
	}

	// // 2. 获取*gdb.Model对象
	var (
		md = dao.Stock.Ctx(ctx)
	)

	md.Where(dao.Stock.Columns().DealerId,

	return
}
