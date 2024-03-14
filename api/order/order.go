package order

import (
	"SheeDrive/api/pagination"
	"SheeDrive/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 订单信息查询
type OrderGetListReq struct {
	g.Meta `path:"/list" method:"get"`
	pagination.CommonPaginationReq
	// 关键字查询可选字段
	UserName   string      `p:"user_name" dc:"用户姓名"`
	DealerName string      `p:"dealer_name" dc:"经销商名称"`
	CarName    string      `p:"car_name" dc:"汽车名称"`
	Status     int         `p:"status" d:"-2" dc:"订单状态"`
	OrderDate  *gtime.Time `p:"order_date" v:"date#请输入正确的日期格式" dc:"预定日期"`
}

type OrderGetListRes struct {
	pagination.CommonPaginationRes
}

// 通过Id查询订单
type OrderGetByIdReq struct {
	g.Meta `path:"/detail" method:"get"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type OrderGetByIdRes struct {
	OrderInfo model.OrderInfoBase `json:"orderInfo" dc:"订单信息"`
}

type OrderAddReq struct {
	g.Meta    `path:"/add" method:"post"`
	UserId    int64       `p:"userId" v:"required#请输入用户Id" dc:"用户Id"`
	DealerId  int64       `p:"dealerId" v:"required#请输入经销商Id" dc:"经销商Id"`
	CarId     int64       `p:"carId" v:"required#请输入汽车Id" dc:"汽车Id"`
	AddrId    int64       `p:"addrId" v:"required#请输入地址Id" dc:"地址Id"`
	OrderTime *gtime.Time `p:"orderTime" v:"required|date#请选择预定时间|请输入正确的日期格式" dc:"预定时间"`
}

type OrderAddRes struct {
	OrderInfo model.OrderInfoBase `json:"orderInfo" dc:"订单信息"`
}

type OrderDeleteReq struct {
	g.Meta `path:"/delete" method:"delete"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type OrderDeleteRes struct{}

type OrderUpdateCancelReq struct {
	g.Meta `path:"/update/cancel" method:"put"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type OrderUpdateCancelRes struct{}

type OrderUpdateConfirmReq struct {
	g.Meta `path:"/update/confirm" method:"put"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type OrderUpdateConfirmRes struct{}

type OrderUpdateSignReq struct {
	g.Meta `path:"/update/sign" method:"put"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type OrderUpdateSignRes struct{}

type OrderUpdateStartReq struct {
	g.Meta `path:"/update/start" method:"put"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type OrderUpdateStartRes struct{}

type OrderUpdateEndReq struct {
	g.Meta `path:"/update/end" method:"put"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type OrderUpdateEndRes struct{}

type OrderUpdateEndAllReq struct {
	g.Meta `path:"/update/endAll" method:"put"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type OrderUpdateEndAllRes struct{}

// 获取汽车排行
type OrderGetCarRankReq struct {
	g.Meta `path:"/get/carRank" method:"get"`
}

type OrderGetCarRankRes struct {
	pagination.CommonPaginationRes
}

type OrderGetIncompleteReq struct {
	g.Meta   `path:"/get/incomplete" method:"get"`
	DealerId int64 `p:"dealer_id" v:"required#请输入经销商ID" dc:"经销商名称"`
}

type OrderGetIncompleteRes struct {
	Total int `json:"total" dc:"未完成订单数"`
}

type OrderGetAddressTimesReq struct {
	g.Meta   `path:"/get/address/times" method:"get"`
	DealerId int64 `p:"dealer_id" v:"required#请输入经销商ID" dc:"经销商名称"`
}

type OrderGetAddressTimesRes struct {
	pagination.CommonPaginationRes
}

type OrderGetTimeCountReq struct {
	g.Meta   `path:"/get/times" method:"get"`
	DealerId int64 `p:"dealer_id" v:"required#请输入经销商ID" dc:"经销商名称"`
}

type OrderGetTimeCountRes struct {
	// 切片有顺序
	TimeSeries []string `json:"timeSeries"`
	OrderCount []int    `json:"orderCount"`
}

type OrderGetByUserIdReq struct {
	g.Meta `path:"/getByUserId" method:"get"`
	pagination.CommonPaginationReq
	UserId int64 `p:"userId" v:"required#请输入用户ID" dc:"用户ID"`
}

type OrderGetByUserIdRes struct {
	pagination.CommonPaginationRes
}
