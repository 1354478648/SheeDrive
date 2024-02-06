// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Article is the golang structure of table article for DAO operations like Where/Data.
type Article struct {
	g.Meta     `orm:"table:article, do:true"`
	Id         interface{} // 主键ID
	BelongId   interface{} // 作者ID
	IsTop      interface{} // 是否置顶 0否 1是
	Title      interface{} // 标题
	Content    interface{} // 内容
	CarId      interface{} // 汽车ID
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeleteTime *gtime.Time // 删除时间
}
