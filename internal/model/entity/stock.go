// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Stock is the golang structure for table stock.
type Stock struct {
	Id         int64       `json:"id"         ` // 主键ID
	DealerId   int64       `json:"dealerId"   ` // 经销商ID
	CarId      int64       `json:"carId"      ` // 车辆ID
	Stock      int         `json:"stock"      ` // 库存数
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
	DeleteTime *gtime.Time `json:"deleteTime" ` // 删除时间
	CreateUser int64       `json:"createUser" ` // 创建人
	UpdateUser int64       `json:"updateUser" ` // 修改人
	DeleteUser int64       `json:"deleteUser" ` // 删除人
}
