// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Article is the golang structure for table article.
type Article struct {
	Id         int64       `json:"id"         ` // 主键ID
	BelongId   int64       `json:"belongId"   ` // 作者ID
	IsTop      int         `json:"isTop"      ` // 是否置顶 0否 1是
	Title      string      `json:"title"      ` // 标题
	Content    string      `json:"content"    ` // 内容
	CarId      int64       `json:"carId"      ` // 汽车ID
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
	DeleteTime *gtime.Time `json:"deleteTime" ` // 删除时间
}
