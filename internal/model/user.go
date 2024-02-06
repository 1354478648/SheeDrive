package model

import "github.com/gogf/gf/v2/os/gtime"

// 用户信息基类
type UserInfoBase struct {
	Id        int64  `json:"id"         ` // 主键ID
	LastName  string `json:"lastName"   ` // 姓
	FirstName string `json:"firstName"  ` // 名
	Username  string `json:"username"   ` // 用户名
	// Password   string      `json:"password"   ` // 密码
	Avatar     string      `json:"avatar"     ` // 头像
	Phone      string      `json:"phone"      ` // 手机号
	IdNumber   string      `json:"idNumber"   ` // 身份证号
	Sex        string      `json:"sex"        ` // 性别
	Birthday   *gtime.Time `json:"birthday"   ` // 生日
	Status     int         `json:"status"     ` // 状态 0:禁用, 1:正常
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
	DeleteTime *gtime.Time `json:"deleteTime" ` // 删除时间
}

// 用户登录
type UserLoginInput struct {
	Username string
	Password string
}

type UserLoginOutput struct {
	UserInfoBase
}

// 用户注册
type UserRegisterInput struct {
	LastName  string
	FirstName string
	IdNumber  string
	Phone     string
	Password  string
}

type UserRegisterOutput struct {
	UserInfoBase
}

type UserGetListInput struct {
	Page       int
	PageSize   int
	Username   string
	Name       string
	Status     int
	BeforeDate *gtime.Time
	AfterDate  *gtime.Time
}

type UserGetListOutput struct {
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
	Total    int            `json:"total"`
	Items    []UserInfoBase `json:"items"`
}

type UserGetByIdInput struct {
	Id int64
}

type UserGetByIdOutput struct {
	UserInfoBase
}

type UserDeleteInput struct {
	Id int64
}

type UserUpdateStatusInput struct {
	Id int64
}

type UserUpdatePasswordInput struct {
	Id          int64
	OldPassword string
	NewPassword string
}

type UserUpdateAvatarInput struct {
	Id  int64
	Url string
}
