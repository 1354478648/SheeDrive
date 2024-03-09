package order

import (
	"SheeDrive/api/pagination"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 订单信息查询
type OrderGetListReq struct {
	g.Meta `path:"/list" method:"get"`
	pagination.CommonPaginationReq
	// 关键字查询可选字段
	UserName   string      `json:"user_name" dc:"用户姓名"`
	DealerName string      `json:"dealer_name" dc:"经销商名称"`
	CarName    string      `json:"car_name" dc:"汽车名称"`
	Status     int         `json:"status" d:"-2" dc:"订单状态"`
	OrderDate  *gtime.Time `json:"order_date" v:"date#请输入正确的日期格式" dc:"预定日期"`
}

type OrderGetListRes struct {
	pagination.CommonPaginationRes
}
