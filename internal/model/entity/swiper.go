// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Swiper is the golang structure for table swiper.
type Swiper struct {
	Id           int64       `json:"id"           ` // 主键ID
	CarId        int64       `json:"carId"        ` // 车辆ID
	ImageUrl     string      `json:"imageUrl"     ` // 图片地址
	DescribeInfo string      `json:"describeInfo" ` // 描述信息
	CreateTime   *gtime.Time `json:"createTime"   ` // 创建时间
	DeleteTime   *gtime.Time `json:"deleteTime"   ` // 删除时间
}
