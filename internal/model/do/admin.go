// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure of table admin for DAO operations like Where/Data.
type Admin struct {
	g.Meta     `orm:"table:admin, do:true"`
	Id         interface{} // 主键ID
	Name       interface{} // 姓名
	Username   interface{} // 用户名
	Password   interface{} // 密码
	Avatar     interface{} // 头像
	Phone      interface{} // 手机号
	Status     interface{} // 状态 0:禁用, 1:正常
	IsRoot     interface{} // 是否是超级管理员 0:否, 1:是
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeleteTime *gtime.Time // 删除时间
}
