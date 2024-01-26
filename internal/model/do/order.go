// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Order is the golang structure of table order for DAO operations like Where/Data.
type Order struct {
	g.Meta     `orm:"table:order, do:true"`
	Id         interface{} // 主键ID
	UserId     interface{} // 用户ID
	DealerId   interface{} // 经销商ID
	CarId      interface{} // 车辆ID
	AddrId     interface{} // 用户地址ID
	Status     interface{} // 订单状态 -1:异常,0:取消,1:未确认,2:已确认,3:签署协议,4:试驾中,5:试驾结束,6:待评价,7:已评价
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeleteTime *gtime.Time // 删除时间
	CreateUser interface{} // 创建人
	UpdateUser interface{} // 修改人
	DeleteUser interface{} // 删除人
}
