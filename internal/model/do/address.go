// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Address is the golang structure of table address for DAO operations like Where/Data.
type Address struct {
	g.Meta         `orm:"table:address, do:true"`
	Id             interface{} // 主键ID
	BelongId       interface{} // 所属ID
	BelongCategory interface{} // 所属分类 1:经销商,2:用户
	LngLat         interface{} // 经纬度
	Province       interface{} // 省
	City           interface{} // 市
	District       interface{} // 区
	Detail         interface{} // 详细地址
	CreateTime     *gtime.Time // 创建时间
	UpdateTime     *gtime.Time // 更新时间
	DeleteTime     *gtime.Time // 删除时间
}
