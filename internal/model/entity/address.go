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
	Latitude       float64     `json:"latitude"       ` // 纬度
	Longitude      float64     `json:"longitude"      ` // 经度
	Accuracy       float64     `json:"accuracy"       ` // 精度
	Country        string      `json:"country"        ` // 国家
	Province       string      `json:"province"       ` // 省
	City           string      `json:"city"           ` // 市
	District       string      `json:"district"       ` // 区
	Street         string      `json:"street"         ` // 街道
	StreetNumber   string      `json:"streetNumber"   ` // 门牌号
	PoiName        string      `json:"poiName"        ` // POI信息
	PostalCode     string      `json:"postalCode"     ` // 邮政编码
	CityCode       string      `json:"cityCode"       ` // 城市代码
	CreateTime     *gtime.Time `json:"createTime"     ` // 创建时间
	UpdateTime     *gtime.Time `json:"updateTime"     ` // 更新时间
	DeleteTime     *gtime.Time `json:"deleteTime"     ` // 删除时间
}
