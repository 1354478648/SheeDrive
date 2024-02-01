// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Address is the golang structure for table address.
type Address struct {
	Id             int64       `json:"id"             ` // 主键ID
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
