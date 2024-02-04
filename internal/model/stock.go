package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type StockInfoBase struct {
	Id         int64       `json:"id"        `
	DealerId   int64       `json:"dealerId"  `
	CarId      int64       `json:"carId"     `
	CreateTime *gtime.Time `json:"createTime"`
	UpdateTime *gtime.Time `json:"updateTime"`
	DeleteTime *gtime.Time `json:"deleteTime"`

	DealerInfo    *DealerInfoBase    `orm:"with:id=dealer_id" json:"dealerInfo"`
	CarDetailInfo *CarDetailInfoBase `orm:"with:id=car_id" json:"carDetailInfo"`
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