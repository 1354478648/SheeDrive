// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Comment is the golang structure for table comment.
type Comment struct {
	Id          int64       `json:"id"          ` // 主键ID
	OrderId     int64       `json:"orderId"     ` // 订单ID
	Content     string      `json:"content"     ` // 评价内容
	TotalScore  int         `json:"totalScore"  ` // 总评分 1~5星
	DealerScore int         `json:"dealerScore" ` // 经销商评分 1~5星
	CarScore    int         `json:"carScore"    ` // 汽车评分 1~5星
	CreateTime  *gtime.Time `json:"createTime"  ` // 创建时间
	DeleteTime  *gtime.Time `json:"deleteTime"  ` // 删除时间
	CreateUser  int64       `json:"createUser"  ` // 创建人
	DeleteUser  int64       `json:"deleteUser"  ` // 删除人
}
