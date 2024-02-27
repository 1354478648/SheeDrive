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
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
}
