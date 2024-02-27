package stock

import (
	"SheeDrive/api/pagination"
	"SheeDrive/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 库存列表分页与关键字查询（管理员）
type StockGetListReq struct {
	g.Meta `path:"/list" method:"get"`
	pagination.CommonPaginationReq
	//关键字查询可选字段
	DealerName string `json:"dealer_name" dc:"经销商名称"`
	CarName    string `json:"car_name" dc:"汽车名称"`
}

type StockGetListRes struct {
	pagination.CommonPaginationRes
}

// 通过Id查询库存
type StockGetByIdReq struct {
	g.Meta `path:"/detail" method:"get"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type StockGetByIdRes struct {
	StockInfo model.StockInfoBase `json:"stock_info" dc:"库存信息"`
}

// 添加库存
type StockAddReq struct {
	g.Meta   `path:"/add" method:"post"`
	DealerId int64 `p:"dealerId" v:"required#请输入经销商Id" dc:"经销商Id"`
	CarId    int64 `p:"carId" v:"required#请输入汽车信息Id" dc:"汽车信息Id"`
}

type StockAddRes struct {
	Id int64 `json:"id" dc:"主键id"`
}

type StockDeleteReq struct {
	g.Meta `path:"/delete" method:"delete"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type StockDeleteRes struct{}
