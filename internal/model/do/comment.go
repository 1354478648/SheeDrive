// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Comment is the golang structure of table comment for DAO operations like Where/Data.
type Comment struct {
	g.Meta      `orm:"table:comment, do:true"`
	Id          interface{} // 主键ID
	OrderId     interface{} // 订单ID
	Content     interface{} // 评价内容
	TotalScore  interface{} // 总评分 1~5星
	DealerScore interface{} // 经销商评分 1~5星
	CarScore    interface{} // 汽车评分 1~5星
	CreateTime  *gtime.Time // 创建时间
	DeleteTime  *gtime.Time // 删除时间
	CreateUser  interface{} // 创建人
	DeleteUser  interface{} // 删除人
}
