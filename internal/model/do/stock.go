// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Stock is the golang structure of table stock for DAO operations like Where/Data.
type Stock struct {
	g.Meta     `orm:"table:stock, do:true"`
	Id         interface{} // 主键ID
	DealerId   interface{} // 经销商ID
	CarId      interface{} // 车辆ID
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeleteTime *gtime.Time // 删除时间
}
