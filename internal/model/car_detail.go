package model

import "SheeDrive/internal/model/entity"

// 汽车信息添加修改基类
type CarDetailAddUpdateBase struct {
	Year         string
	Brand        string
	Model        string
	Version      string
	Image        string
	Category     int
	Color        string
	Price        int64
	Type         int
	Seats        int
	DescribeInfo string
}

type CarDetailGetListInput struct {
	Page      int
	PageSize  int
	Year      string
	Brand     string
	Model     string
	Category  string
	LowPrice  int64
	HighPrice int64
}

type CarDetailGetListOutput struct {
	Page     int                `json:"page"`
	PageSize int                `json:"page_size"`
	Total    int                `json:"total"`
	Items    []entity.CarDetail `json:"items"`
}

type CarDetailGetByIdInput struct {
	Id int64
}

type CarDetailGetByIdOutput struct {
	CarDetail entity.CarDetail `json:"car_detail"`
}

type CarDetailAddInput struct {
	CarDetailAddUpdateBase
}

type CarDetailAddOutput struct {
	Id int64 `json:"id"`
}

type CarDetailUpdateInput struct {
	Id int64
	CarDetailAddUpdateBase
}

type CarDetailDeleteInput struct {
	Id int64
}
