// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Swiper is the golang structure of table swiper for DAO operations like Where/Data.
type Swiper struct {
	g.Meta       `orm:"table:swiper, do:true"`
	Id           interface{} // 主键ID
	CarId        interface{} // 车辆ID
	ImageUrl     interface{} // 图片地址
	DescribeInfo interface{} // 描述信息
	CreateTime   *gtime.Time // 创建时间
	DeleteTime   *gtime.Time // 删除时间
}
