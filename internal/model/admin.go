package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type AdminBase struct {
	Id       int64  `json:"id"         ` // 主键ID
	Name     string `json:"name"       ` // 姓名
	Username string `json:"username"   ` // 用户名
	// Password   string      `json:"password"   ` // 密码
	Avatar     string      `json:"avatar"     ` // 头像
	Phone      string      `json:"phone"      ` // 手机号
	Status     int         `json:"status"     ` // 状态 0:禁用, 1:正常
	IsRoot     int         `json:"isRoot"     ` // 是否是超级管理员 0:否, 1:是
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
	DeleteTime *gtime.Time `json:"deleteTime" ` // 删除时间
}

// AdminLoginInput 管理员登录
type AdminLoginInput struct {
	Username string
	Password string
}

// AdminLoginOutput 管理员登录结果
type AdminLoginOutput struct {
	AdminBase
}

// AdminGetListInput 分页与关键字查询管理员列表
type AdminGetListInput struct {
	Page       int
	PageSize   int
	Username   string
	Name       string
	BeforeDate *gtime.Time
	AfterDate  *gtime.Time
}

// AdminGetListOutput 分页与关键字查询管理员列表结果
type AdminGetListOutput struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Total    int         `json:"total"`
	Items    []AdminBase `json:"items"`
}
