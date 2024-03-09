package model

import "github.com/gogf/gf/v2/os/gtime"

type AddressInfoBase struct {
	Id             string      `json:"id"             ` // 主键ID
	BelongId       int64       `json:"belongId"       ` // 所属ID
	BelongCategory int         `json:"belongCategory" ` // 所属分类 1:经销商,2:用户
	LngLat         string      `json:"lngLat"         ` // 经纬度
	Province       string      `json:"province"       ` // 省
	City           string      `json:"city"           ` // 市
	District       string      `json:"district"       ` // 区
	Detail         string      `json:"detail"         ` // 详细地址
	CreateTime     *gtime.Time `json:"createTime"     ` // 创建时间
	UpdateTime     *gtime.Time `json:"updateTime"     ` // 更新时间
	DeleteTime     *gtime.Time `json:"deleteTime"     ` // 删除时间
}

// 用户地址添加修改基类
type UserAddressAddUpdateBase struct {
	Province string `json:"province"       ` // 省
	City     string `json:"city"           ` // 市
	District string `json:"district"       ` // 区
	Detail   string `json:"detail"         ` // 详细地址

}

// 用户地址列表查询
type UserAddressGetListInput struct {
	BelongId       int64
	BelongCategory int
}

type UserAddressGetListOutput struct {
	Items []AddressInfoBase `json:"items"`
}

// 用户地址添加
type UserAddressAddInput struct {
	BelongId int64
	UserAddressAddUpdateBase
}

type UserAddressAddOutput struct {
	Id int64
}

type UserAddressGetByIdInput struct {
	Id int64
}

type UserAddressGetByIdOutput struct {
	AddressInfoBase
}

type UserAddressUpdateInput struct {
	Id int64
	UserAddressAddUpdateBase
}

type UserAddressDeleteInput struct {
	Id int64
}
