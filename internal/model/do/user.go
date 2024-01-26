// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta     `orm:"table:user, do:true"`
	Id         interface{} // 主键ID
	LastName   interface{} // 姓
	FirstName  interface{} // 名
	Username   interface{} // 用户名
	Password   interface{} // 密码
	Avatar     interface{} // 头像
	Phone      interface{} // 手机号
	IdNumber   interface{} // 身份证号
	Sex        interface{} // 性别
	Birthday   *gtime.Time // 生日
	Status     interface{} // 状态 0:禁用, 1:正常
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	DeleteTime *gtime.Time // 删除时间
	CreateUser interface{} // 创建人
	UpdateUser interface{} // 修改人
	DeleteUser interface{} // 删除人
}
