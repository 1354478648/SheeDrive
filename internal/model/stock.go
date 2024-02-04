package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type StockInfo struct {
	g.Meta     `orm:"table:stock"`
	Id         int64       `json:"id"         orm:"id"`
	DealerId   int64       `json:"dealerId"   orm:"dealer_id"`
	CarId      int64       `json:"carId"      orm:"car_id"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time"`
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time"`
	DeleteTime *gtime.Time `json:"deleteTime" orm:"delete_time"`
}

// 库存信息基类
type StockInfoBase struct {
	StockInfo *StockInfo
	Dealer    *DealerInfoBase
	CarDetail *CarDetailInfoBase
}

type StockGetListInput struct {
	Page       int
	PageSize   int
	DealerName string
}

type StockGetListOutput struct {
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
	Total    int             `json:"total"`
	Items    []StockInfoBase `json:"items"`
}
