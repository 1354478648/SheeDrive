// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Order is the golang structure for table order.
type Order struct {
	Id         int64       `json:"id"         ` // 主键ID
	UserId     int64       `json:"userId"     ` // 用户ID
	DealerId   int64       `json:"dealerId"   ` // 经销商ID
	CarId      int64       `json:"carId"      ` // 车辆ID
	AddrId     int64       `json:"addrId"     ` // 用户地址ID
	Status     int         `json:"status"     ` // 订单状态 -1:异常,0:取消,1:未确认,2:已确认,3:签署协议,4:试驾中,5:试驾结束,6:待评价,7:已评价
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
	DeleteTime *gtime.Time `json:"deleteTime" ` // 删除时间
	CreateUser int64       `json:"createUser" ` // 创建人
	UpdateUser int64       `json:"updateUser" ` // 修改人
	DeleteUser int64       `json:"deleteUser" ` // 删除人
}
