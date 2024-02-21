// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id         int64       `json:"id"         ` // 主键ID
	LastName   string      `json:"lastName"   ` // 姓
	FirstName  string      `json:"firstName"  ` // 名
	Username   string      `json:"username"   ` // 用户名
	Password   string      `json:"password"   ` // 密码
	Avatar     string      `json:"avatar"     ` // 头像
	Phone      string      `json:"phone"      ` // 手机号
	IdNumber   string      `json:"idNumber"   ` // 身份证号
	Sex        string      `json:"sex"        ` // 性别
	Birthday   *gtime.Time `json:"birthday"   ` // 生日
	Status     int         `json:"status"     ` // 状态 0:禁用, 1:正常
	Token      string      `json:"token"      ` // token
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
	DeleteTime *gtime.Time `json:"deleteTime" ` // 删除时间
}
