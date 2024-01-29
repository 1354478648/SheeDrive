package model

import "github.com/gogf/gf/v2/os/gtime"

// 经销商信息基类
type DealerInfoBase struct {
	Id       int64  `json:"id"         ` // 主键ID
	Name     string `json:"name"       ` // 姓名
	Username string `json:"username"   ` // 用户名
	// Password   string      `json:"password"   ` // 密码
	Avatar       string      `json:"avatar"     ` // 头像
	Phone        string      `json:"phone"      ` // 手机号
	DescribeInfo string      `json:"describeInfo"`
	Status       int         `json:"status"     ` // 状态 0:禁用, 1:正常
	CreateTime   *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime   *gtime.Time `json:"updateTime" ` // 更新时间
	DeleteTime   *gtime.Time `json:"deleteTime" ` // 删除时间
}

// 经销商添加修改基类
type DealerAddUpdateBase struct {
	Name         string `json:"name"`     //姓名
	Username     string `json:"username"` //用户名
	Phone        string `json:"phone"`    //手机号
	DescribeInfo string `json:"describeInfo"`
}

type DealerLoginInput struct {
	Username string
	Password string
}

type DealerLoginOutput struct {
	DealerInfoBase
}
